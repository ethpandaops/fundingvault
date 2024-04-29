import {
	useAccount,
  useReadContract,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import { useEffect, useState } from "react";
import { toBigintUnit, toDecimalUnit, toReadableAmount, toReadableDuration } from "../../utils/ConvertHelpers";
import { isAddress } from "ethers";

import "./ClaimForm.css"

interface IGrantDetails {
  claimInterval: bigint
  claimLimit: bigint
  claimTime: bigint
  dustBalance: bigint
}

const ClaimForm = (props: { grantId: number }): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [claimAmount, setClaimAmount] = useState("10");
  let [claimTarget, setClaimTarget] = useState("");
  let [claimAll, setClaimAll] = useState<boolean>(false);
  let [claimTargetCustom, setClaimTargetCustom] = useState<boolean>(false);

  const grantDetails = useReadContract({
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getGrant",
		args: [ props.grantId ],
	});
  const claimableBalance = useReadContract({
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getClaimableBalance",
		args: [ props.grantId ],
	});
  const claimRequest = useWriteContract();

  //console.log(grantDetails.data);
  useEffect(() => {
    const interval = setInterval(() => {
      console.log("refetch");
      claimableBalance.refetch();
    }, 15000);
    return () => {
      clearInterval(interval);
    };
  }, []);
  
  let maxAmount = toDecimalUnit(claimableBalance.data as bigint, chain?.nativeCurrency.decimals);
  if(isNaN(maxAmount)) {
    maxAmount = 0;
  }
  maxAmount = Math.round(maxAmount * 1000) / 1000;

  if(parseInt(claimAmount) > maxAmount) {
    setClaimAmount(maxAmount.toString());
  } else if(parseInt(claimAmount) < 0) {
    setClaimAmount("0");
  }

  return (
    <div>
      <table className="details-table">
        <tbody>
          <tr>
            <td className="prop">Your claimable balance:</td>
            <td className="value">{toReadableAmount(claimableBalance.data as bigint, chain?.nativeCurrency.decimals, chainConfig.TokenName, 3)}</td>
          </tr>
          {grantDetails.data ?
          <tr>
            <td className="prop">Your allowance:</td>
            <td className="value">{toReadableAmount((grantDetails.data as IGrantDetails)?.claimLimit, 0, chainConfig.TokenName, 0)} per {toReadableDuration((grantDetails.data as IGrantDetails)?.claimInterval)}</td>
          </tr>
          : null}
        </tbody>
      </table>
      <div className="claim-form container">
        <b>Claim Funds</b>

        <div className="row mt-2">
          <div className="col-6">
            Amount ({chainConfig.TokenName})
          </div>
          <div className="col-6">
            <div className="form-check">
              <input className="form-check-input" type="checkbox" value="" id="claimAll" onChange={(evt) => setClaimAll(evt.target.checked)} checked={claimAll} />
              <label className="form-check-label" htmlFor="claimAll">
                Claim all claimable balance
              </label>
            </div>
          </div>
        </div>
        <div className="row">
          <div className="col-5">
            <input type="number" className="form-control" placeholder={claimAll ? maxAmount.toString() : "0"} onChange={(evt) => setClaimAmount(evt.target.value)} value={claimAll ? maxAmount.toString() : claimAmount} disabled={claimAll} />
          </div>
          <div className="col-1"></div>
          <div className="col-6">
            <input type="range" className="form-range" max={maxAmount} onChange={(evt) => setClaimAmount(evt.target.value)} value={claimAll ? maxAmount.toString() : claimAmount} disabled={claimAll} />
          </div>
        </div>

        <div className="row mt-2">
          <div className="col-6">
            Target Wallet
          </div>
          <div className="col-6">
            <div className="form-check">
              <input className="form-check-input" type="checkbox" value="" id="claimTargetCustom" onChange={(evt) => setClaimTargetCustom(evt.target.checked)} checked={claimTargetCustom} />
              <label className="form-check-label" htmlFor="claimTargetCustom">
                Send to another wallet
              </label>
            </div>
          </div>
        </div>
        <div className="row">
          <div className="col-12">
            <input type="text" className="form-control" placeholder={claimTargetCustom ? "0x..." : address} onChange={(evt) => setClaimTarget(evt.target.value)} value={claimTarget} disabled={!claimTargetCustom} />
          </div>
        </div>

        <div className="row mt-3">
          <div className="col-12">
            <button className="btn btn-primary claim-button" onClick={(evt) => requestFunds(evt.target as HTMLButtonElement)} disabled={claimRequest.isPending}>Request Funds</button>
          </div>
        </div>

        {claimRequest.isPending && claimRequest.data as any ?
        <div className="row mt-3">
          <div className="col-12">
            <div className="alert alert-info">
              Claim transaction pending... TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + claimRequest.data} target="_blank">{claimRequest.data}</a>
            </div>
          </div>
        </div>
        : null}
        {claimRequest.isError ?
        <div className="row mt-3">
          <div className="col-12">
            <div className="alert alert-danger">
              Claim failed. {claimRequest.data as any ? <span>TX: <a href={chainConfig.BlockExplorerUrl + "tx/" + claimRequest.data} target="_blank">{claimRequest.data}</a></span> : null}<br />
              {claimRequest.error.message}
            </div>
          </div>
        </div>
        : null}
        {claimRequest.isSuccess ?
        <div className="row mt-3">
          <div className="col-12">
            <div className="alert alert-success">
              Claim TX: <a className="txhash" href={chainConfig.BlockExplorerUrl + "tx/" + claimRequest.data} target="_blank">{claimRequest.data}</a>
            </div>
          </div>
        </div>
        : null}
      </div>

    </div>
  )

  function requestFunds(button: HTMLButtonElement) {
    button.disabled = true;

    let targetAddress = claimTarget;
    if(claimTargetCustom && !isAddress(targetAddress)) {
      alert("Provided target address '" + targetAddress + "' is invalid.");
      button.disabled = false;
      return;
    }

    let amount = parseInt(claimAmount);
    if(claimAll) {
      amount = 0;
    } else if(amount == 0 || amount > maxAmount) {
      alert("Desired amount '" + claimAmount + "' is invalid.");
      button.disabled = false;
      return;
    }
    let amountWei = toBigintUnit(amount, chain?.nativeCurrency.decimals);

    let callfn = "claim"
    let callArgs: any[] = [ props.grantId, amountWei ];
    if (claimTargetCustom && targetAddress.toLowerCase() != address?.toLowerCase()) {
      callfn = "claimTo";
      callArgs.push(targetAddress);
    }

    claimRequest.writeContract({
      address: chainConfig.VaultContractAddr,
      account: address,
      abi: FundingVaultAbi,
      chainId: chainConfig.Chain.id,
      functionName: callfn,
      args: callArgs,
    })
  }
}

export default ClaimForm;