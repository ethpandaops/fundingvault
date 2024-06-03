import {
	useAccount,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import { Modal } from 'react-bootstrap';

const GrantDelete = (props: { grantId: number, name: string, owner: string, closeFn?: () => void }): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  
  const deleteRequest = useWriteContract();

  return (
    <Modal show centered className="grant-manager-dialog delete-dialog" size="lg" onHide={() => {
      if(props.closeFn)
        props.closeFn();
    }}>
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Delete Grant
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
              Owner:
            </div>
            <div className="col-8">
              <a href={chainConfig.BlockExplorerUrl + "address/" + props.owner} target="_blank" rel="noreferrer">{props.owner}</a>
            </div>
          </div>

          {deleteRequest.isPending && deleteRequest.data as any ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-info">
                Delete transaction pending... TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + deleteRequest.data} target="_blank" rel="noreferrer">{deleteRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
          {deleteRequest.isError ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-danger">
                Delete failed. {deleteRequest.data as any ? <span>TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + deleteRequest.data} target="_blank" rel="noreferrer">{deleteRequest.data}</a></span> : null}<br />
                {deleteRequest.error.message}
              </div>
            </div>
          </div>
          : null}
          {deleteRequest.isSuccess ?
          <div className="row mt-3">
            <div className="col-12">
              <div className="alert alert-success">
                Delete TX: <a className="txhash" href={chainConfig.BlockExplorerUrl + "tx/" + deleteRequest.data} target="_blank" rel="noreferrer">{deleteRequest.data}</a>
              </div>
            </div>
          </div>
          : null}
        </div>
      </Modal.Body>
      <Modal.Footer>
        <button onClick={(evt) => removeGrant(evt.target as HTMLButtonElement)} disabled={deleteRequest.isPending} className="btn btn-primary">
          Delete Grant
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

  function removeGrant(button: HTMLButtonElement) {
    button.disabled = true;

    deleteRequest.writeContract({
      address: chainConfig.VaultContractAddr,
      account: address,
      abi: FundingVaultAbi,
      chainId: chainConfig.Chain.id,
      functionName: "removeGrant",
      args: [ props.grantId ],
      onSuccess: () => {
      },
    })
  }
  
}

export default GrantDelete;