import {
	useAccount,
} from "wagmi";
import {
  useConnectModal,
} from '@rainbow-me/rainbowkit';
import EligibilityCheck from "./EligibilityCheck";


const VaultInfo = (): React.ReactElement => {
  const { isConnected, chain } = useAccount();
  const { openConnectModal } = useConnectModal();

  return (
    <div className="page-block">
      <h1>Testnet Funding Vault</h1>
      <p>
        The FundingVault contract provides a way to distribute continuous limited amounts of funds to authorized entities.<br />
        The distribution is time gated and a specific limit per grant is enforced.<br />
        Check out the <a href="https://github.com/ethpandaops/fundingvault/blob/master/README.md">FundingVault repository</a> for more details.
      </p>
      {isConnected && chain ?
      <EligibilityCheck /> : null}
      {!isConnected ? renderDisconnected() : null}
      {isConnected && !chain ?  renderInvalidNetwork() : null}
    </div>
  )

  function renderDisconnected() {
    return (
      <div className="">
        Please <a href="#" onClick={openConnectModal}>connect to your wallet</a> to continue.
      </div>
    )
  }

  function renderInvalidNetwork() {
    return (
      <div className="">
        Please switch to holesky or sepolia to continue.
      </div>
    )
  }
}

export default VaultInfo;