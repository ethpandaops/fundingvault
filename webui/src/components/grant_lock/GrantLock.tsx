import {
	useAccount,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import { useState } from "react";
import { Modal } from 'react-bootstrap';

const GrantLock = (props: { grantId: number, name: string, closeFn?: () => void }): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [timeInput, setTimeInput] = useState(0);
  
  const lockRequest = useWriteContract();

  return (
    <Modal show centered className="grant-manager-dialog rename-dialog" size="lg" onHide={() => {
      if(props.closeFn)
        props.closeFn();
    }}>
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Lock Grant
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
              Lock Time:
            </div>
            <div className="col-8">
              <input type="number" className="form-control" onChange={(evt) => setTimeInput(parseInt(evt.target.value))} value={timeInput} />
            </div>
          </div>

          {lockRequest.isPending && lockRequest.data as any ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-info">
                Lock transaction pending... TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + lockRequest.data} target="_blank" rel="noreferrer">{lockRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
          {lockRequest.isError ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-danger">
                Lock failed. {lockRequest.data as any ? <span>TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + lockRequest.data} target="_blank" rel="noreferrer">{lockRequest.data}</a></span> : null}<br />
                {lockRequest.error.message}
              </div>
            </div>
          </div>
          : null}
          {lockRequest.isSuccess ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-success">
                Lock TX: <a className="txhash" href={chainConfig.BlockExplorerUrl + "tx/" + lockRequest.data} target="_blank" rel="noreferrer">{lockRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
        </div>
      </Modal.Body>
      <Modal.Footer>
        <button onClick={(evt) => lockGrant(evt.target as HTMLButtonElement)} disabled={lockRequest.isPending} className="btn btn-primary">
          Lock Grant
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

  function lockGrant(button: HTMLButtonElement) {
    button.disabled = true;

    lockRequest.writeContract({
      address: chainConfig.VaultContractAddr,
      account: address,
      abi: FundingVaultAbi,
      chainId: chainConfig.Chain.id,
      functionName: "lockGrant",
      args: [ props.grantId, timeInput ],
      onSuccess: () => {},
    })
  }
  
}

export default GrantLock;