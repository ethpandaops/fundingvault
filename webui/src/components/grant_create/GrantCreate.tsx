import {
	useAccount,
  useReadContract,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import { useState } from "react";
import { Modal } from 'react-bootstrap';
import { toReadableAmount, toReadableDuration } from "../../utils/ConvertHelpers";
import { isAddress } from "ethers";

function toHex(str) {
  var result = '';
  var ccode;
  for (var i=0; i<str.length; i++) {
    ccode = str.charCodeAt(i);
    if(ccode) {
      result += ccode.toString(16);
    }
  }
  return result;
}

const GrantCreate = (props: { closeFn?: () => void }): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [nameInput, setNameInput] = useState("");
  let [addressInput, setAddressInput] = useState("");
  let [amountInput, setAmountInput] = useState(10000);
  let [intervalInput, setIntervalInput] = useState(2592000);
  
  const createRequest = useWriteContract();

  let intervalOptions = [
    { value: 86400, title: "1 day" },
    { value: 604800, title: "1 week" },
    { value: 1209600, title: "2 weeks" },
    { value: 2592000, title: "1 month" },
  ];
  let intervalIsOption = false;
  intervalOptions.forEach((option) => {
    if(option.value == intervalInput)
      intervalIsOption = true;
  });

  return (
    <Modal show centered className="grant-manager-dialog create-dialog" size="lg" onHide={() => {
      if(props.closeFn)
        props.closeFn();
    }}>
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Create Grant
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <div className="container">
          <div className="row my-2">
            <div className="col-2">
              Name:
            </div>
            <div className="col-8">
              <input type="text" maxLength={32} className="form-control" onChange={(evt) => setNameInput(evt.target.value)} value={nameInput} />
            </div>
          </div>
          <div className="row my-2">
            <div className="col-2">
              Address:
            </div>
            <div className="col-8">
              <input type="text" className="form-control" placeholder="0x..." onChange={(evt) => setAddressInput(evt.target.value)} value={addressInput} />
            </div>
          </div>
          <div className="row my-2">
            <div className="col-2">
              <span className="field-label">Amount:</span>
            </div>
            <div className="col-4">
              <input type="number" className="form-control" onChange={(evt) => setAmountInput(parseInt(evt.target.value))} value={amountInput.toString()} />
            </div>
            <div className="col-2">
              <span className="field-label">{chainConfig.TokenName}</span>
            </div>
          </div>
          <div className="row my-2">
            <div className="col-2">
              <span className="field-label">Interval:</span>
            </div>
            <div className="col-5">
              <select className="form-select" onChange={(evt) => { setIntervalInput(parseInt(evt.target.value)); }}>
                {intervalOptions.map((option) => {
                  return (
                    <option key={option.value} value={option.value} selected={option.value == intervalInput}>{option.title}</option>
                  );
                })}
                <option value="0" selected={!intervalIsOption}>custom</option>
              </select>
            </div>
          </div>
          {intervalIsOption ? null :
            <div className="row">
              <div className="col-2">
              </div>
              <div className="col-4">
                <input type="number" className="form-control" onChange={(evt) => setIntervalInput(parseInt(evt.target.value))} value={intervalInput.toString()} />
              </div>
              <div className="col-1">
                <span className="field-label">sec</span>
              </div>
            </div>
          }

          {createRequest.isPending && createRequest.data as any ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-info">
                Create transaction pending... TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + createRequest.data} target="_blank">{createRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
          {createRequest.isError ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-danger">
                Create failed. {createRequest.data as any ? <span>TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + createRequest.data} target="_blank">{createRequest.data}</a></span> : null}<br />
                {createRequest.error.message}
              </div>
            </div>
          </div>
          : null}
          {createRequest.isSuccess ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-success">
                Create TX: <a className="txhash" href={chainConfig.BlockExplorerUrl + "tx/" + createRequest.data} target="_blank">{createRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
        </div>
      </Modal.Body>
      <Modal.Footer>
        <button onClick={(evt) => createGrant(evt.target as HTMLButtonElement)} disabled={createRequest.isPending} className="btn btn-primary">
          Create Grant
        </button>
        <button onClick={() => {
          if(props.closeFn)
            props.closeFn();
        }} className="btn btn-secondary">
          Cancel
        </button>
      </Modal.Footer>
    </Modal>
  );

  function createGrant(button: HTMLButtonElement) {
    button.disabled = true;

    if(!isAddress(addressInput)) {
      alert("Provided target address '" + addressInput + "' is invalid.");
      button.disabled = false;
      return;
    }

    let hexName = toHex(nameInput);
    while(hexName.length < 64) {
      hexName += "0";
    }
    hexName = "0x" + hexName;

    createRequest.writeContract({
      address: chainConfig.VaultContractAddr,
      account: address,
      abi: FundingVaultAbi,
      chainId: chainConfig.Chain.id,
      functionName: "createGrant",
      args: [ addressInput, amountInput, intervalInput ],
    })
  }
  
}

export default GrantCreate;