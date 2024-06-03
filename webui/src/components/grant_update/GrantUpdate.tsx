import {
	useAccount,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import { useState } from "react";
import { Modal } from 'react-bootstrap';
import { toReadableAmount, toReadableDuration } from "../../utils/ConvertHelpers";

const GrantRename = (props: { grantId: number, name: string, amount: number, interval: number, closeFn?: () => void }): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [amountInput, setAmountInput] = useState(props.amount);
  let [intervalInput, setIntervalInput] = useState(props.interval);
  
  const updateRequest = useWriteContract();

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
    <Modal show centered className="grant-manager-dialog update-dialog" size="lg" onHide={() => {
      if(props.closeFn)
        props.closeFn();
    }}>
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Update Grant
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <div className="container">
          <div className="row my-2">
            <div className="col-2">
              Grant ID:
            </div>
            <div className="col-8">
              {props.grantId} ({props.name})
            </div>
          </div>
          <div className="row my-2">
            <div className="col-2">
              Allowance:
            </div>
            <div className="col-8">
              {toReadableAmount(props.amount, 0, chainConfig.TokenName, 0)} / {toReadableDuration(props.interval)}
            </div>
          </div>
          <div className="row my-2">
            <div className="col-2">
              <span className="field-label">New Amount:</span>
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
              <span className="field-label">New Interval:</span>
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

          {updateRequest.isPending && updateRequest.data as any ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-info">
                Update transaction pending... TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + updateRequest.data} target="_blank" rel="noreferrer">{updateRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
          {updateRequest.isError ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-danger">
                Update failed. {updateRequest.data as any ? <span>TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + updateRequest.data} target="_blank" rel="noreferrer">{updateRequest.data}</a></span> : null}<br />
                {updateRequest.error.message}
              </div>
            </div>
          </div>
          : null}
          {updateRequest.isSuccess ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-success">
                Rename TX: <a className="txhash" href={chainConfig.BlockExplorerUrl + "tx/" + updateRequest.data} target="_blank" rel="noreferrer">{updateRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
        </div>
      </Modal.Body>
      <Modal.Footer>
        <button onClick={(evt) => updateGrant(evt.target as HTMLButtonElement)} disabled={updateRequest.isPending} className="btn btn-primary">
          Update Grant
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

  function updateGrant(button: HTMLButtonElement) {
    button.disabled = true;

    updateRequest.writeContract({
      address: chainConfig.VaultContractAddr,
      account: address,
      abi: FundingVaultAbi,
      chainId: chainConfig.Chain.id,
      functionName: "updateGrant",
      args: [ props.grantId, amountInput, intervalInput ],
    })
  }
  
}

export default GrantRename;