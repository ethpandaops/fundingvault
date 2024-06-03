import {
	useAccount,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import { useState } from "react";
import { Modal } from 'react-bootstrap';

const GrantTransfer = (props: { grantId: number, name: string, owner: string, closeFn?: () => void }): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [addressInput, setAddressInput] = useState(props.owner);
  
  const transferRequest = useWriteContract();

  return (
    <Modal show centered className="grant-manager-dialog transfer-dialog" size="lg" onHide={() => {
      if(props.closeFn)
        props.closeFn();
    }}>
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Transfer Grant
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <div className="container">
          <div className="row my-2">
            <div className="col-3">
              Grant ID:
            </div>
            <div className="col-7">
              {props.grantId} ({props.name})
            </div>
          </div>
          <div className="row my-2">
            <div className="col-3">
              Current Owner:
            </div>
            <div className="col-7">
              <a href={chainConfig.BlockExplorerUrl + "address/" + props.owner} target="_blank" rel="noreferrer">{props.owner}</a>
            </div>
          </div>
          <div className="row my-2">
            <div className="col-3">
              New Address:
            </div>
            <div className="col-7">
              <input type="text" maxLength={42} className="form-control" onChange={(evt) => setAddressInput(evt.target.value)} value={addressInput} />
            </div>
          </div>

          {transferRequest.isPending && transferRequest.data as any ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-info">
                Transfer transaction pending... TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + transferRequest.data} target="_blank" rel="noreferrer">{transferRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
          {transferRequest.isError ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-danger">
                Transfer failed. {transferRequest.data as any ? <span>TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + transferRequest.data} target="_blank" rel="noreferrer">{transferRequest.data}</a></span> : null}<br />
                {transferRequest.error.message}
              </div>
            </div>
          </div>
          : null}
          {transferRequest.isSuccess ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-success">
                Transfer TX: <a className="txhash" href={chainConfig.BlockExplorerUrl + "tx/" + transferRequest.data} target="_blank" rel="noreferrer">{transferRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
        </div>
      </Modal.Body>
      <Modal.Footer>
        <button onClick={(evt) => renameGrant(evt.target as HTMLButtonElement)} disabled={transferRequest.isPending} className="btn btn-primary">
          Transfer Grant
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

    transferRequest.writeContract({
      address: chainConfig.VaultContractAddr,
      account: address,
      abi: FundingVaultAbi,
      chainId: chainConfig.Chain.id,
      functionName: "transferGrant",
      args: [ props.grantId, addressInput ],
      onSuccess: () => {
        props.name = addressInput;
      },
    })
  }
  
}

export default GrantTransfer;