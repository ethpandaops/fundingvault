import {
	useAccount,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import { useState } from "react";
import { Modal } from 'react-bootstrap';

export interface IGrantDetails {
  claimInterval: bigint
  claimLimit: bigint
  claimTime: bigint
  dustBalance: bigint
}

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

const GrantRename = (props: { grantId: number, name: string, closeFn?: () => void }): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [nameInput, setNameInput] = useState(props.name);
  
  const renameRequest = useWriteContract();

  return (
    <Modal show centered className="grant-manager-dialog rename-dialog" size="lg" onHide={() => {
      if(props.closeFn)
        props.closeFn();
    }}>
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Rename Grant
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
              New Name:
            </div>
            <div className="col-8">
              <input type="text" maxLength={32} className="form-control" onChange={(evt) => setNameInput(evt.target.value)} value={nameInput} />
            </div>
          </div>

          {renameRequest.isPending && renameRequest.data as any ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-info">
                Rename transaction pending... TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + renameRequest.data} target="_blank" rel="noreferrer">{renameRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
          {renameRequest.isError ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-danger">
                Rename failed. {renameRequest.data as any ? <span>TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + renameRequest.data} target="_blank" rel="noreferrer">{renameRequest.data}</a></span> : null}<br />
                {renameRequest.error.message}
              </div>
            </div>
          </div>
          : null}
          {renameRequest.isSuccess ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-success">
                Rename TX: <a className="txhash" href={chainConfig.BlockExplorerUrl + "tx/" + renameRequest.data} target="_blank" rel="noreferrer">{renameRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
        </div>
      </Modal.Body>
      <Modal.Footer>
        <button onClick={(evt) => renameGrant(evt.target as HTMLButtonElement)} disabled={renameRequest.isPending} className="btn btn-primary">
          Rename Grant
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

  function renameGrant(button: HTMLButtonElement) {
    button.disabled = true;

    let hexName = toHex(nameInput);
    while(hexName.length < 64) {
      hexName += "0";
    }
    hexName = "0x" + hexName;
    console.log(hexName);

    renameRequest.writeContract({
      address: chainConfig.VaultContractAddr,
      account: address,
      abi: FundingVaultAbi,
      chainId: chainConfig.Chain.id,
      functionName: "renameGrant",
      args: [ props.grantId, hexName ],
      onSuccess: () => {
        props.name = nameInput;
      },
    })
  }
  
}

export default GrantRename;