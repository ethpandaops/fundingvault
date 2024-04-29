import {
	useAccount,
  useReadContract,
  useWriteContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import VaultTokenAbi from "../../abi/VaultToken.json";
import { useEffect, useState } from "react";
import { toBigintUnit, toDecimalUnit, toReadableAmount, toReadableDuration } from "../../utils/ConvertHelpers";
import { isAddress } from "ethers";
import GrantRename from "../grant_rename/GrantRename";
import GrantUpdate from "../grant_update/GrantUpdate";

export interface IGrantItemProps {
  tokenIdx: number
  grant: IGrantDetails
  setDialog: (element: React.ReactElement) => void;
}

export interface IGrantDetails {
  claimInterval: bigint
  claimLimit: bigint
  claimTime: bigint
  dustBalance: bigint
}

function hex2a(hexx: string): string {
  var hex = hexx.toString();//force conversion
  var str = '';
  if(hex.length >= 2 && hex.substring(0, 2) == "0x")
    hex = hex.substring(2);
  var ccode;
  for (var i = 0; i < hex.length; i += 2) {
    ccode = parseInt(hex.substr(i, 2), 16);
    if(ccode)
      str += String.fromCharCode(ccode);
  }
  return str;
}

const GrantItem = (props: IGrantItemProps): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  
  const tokenIdCall = useReadContract({
		address: chainConfig.TokenContractAddr,
    account: address,
		abi: VaultTokenAbi,
		chainId: chainConfig.Chain.id,
		functionName: "tokenByIndex",
		args: [ props.tokenIdx ],
	});
  const ownerOfCall = useReadContract(tokenIdCall.isFetched ? {
		address: chainConfig.TokenContractAddr,
    account: address,
		abi: VaultTokenAbi,
		chainId: chainConfig.Chain.id,
		functionName: "ownerOf",
		args: [ tokenIdCall.data ],
	} : undefined);
  const claimableBalanceCall = useReadContract(tokenIdCall.isFetched ? {
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getClaimableBalance",
		args: [ tokenIdCall.data ],
	} : undefined);
  const totalClaimedCall = useReadContract(tokenIdCall.isFetched ? {
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getGrantTotalClaimed",
		args: [ tokenIdCall.data ],
	} : undefined);
  const grantNameCall = useReadContract(tokenIdCall.isFetched ? {
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getGrantName",
		args: [ tokenIdCall.data ],
	} : undefined);

  var grantName: string;
  if(grantNameCall.isFetched) {
    grantName = hex2a(grantNameCall.data as string);
  }

  useEffect(() => {
    const interval = setInterval(() => {
      console.log("refetch");
      tokenIdCall.refetch();
      ownerOfCall.refetch();
    }, 15000);
    return () => {
      clearInterval(interval);
    };
  }, []);

  return (
    <tr>
      <td>{tokenIdCall.data?.toString() as string}</td>
      <td><a href={chainConfig.BlockExplorerUrl + "address/" + ownerOfCall.data?.toString()} target="_blank">{ownerOfCall.data?.toString()}</a></td>
      <td>
        {grantName} 
        <a href="#" className="grant-edit-btn" onClick={(evt) => {
          evt.preventDefault();
          if(!tokenIdCall.isFetched)
            return;

          props.setDialog((
            <GrantRename grantId={parseInt(tokenIdCall.data?.toString())} name={grantName} closeFn={() => { props.setDialog(null); }} />
          ));
        }}>
          <i className="bi bi-pencil"></i>
        </a>
      </td>
      <td>
        {toReadableAmount(props.grant.claimLimit as bigint, 0, chainConfig.TokenName, 0)} / {toReadableDuration(props.grant.claimInterval)}
        <a href="#" className="grant-edit-btn" onClick={(evt) => {
          evt.preventDefault();
          if(!tokenIdCall.isFetched)
            return;

          props.setDialog((
            <GrantUpdate 
              grantId={parseInt(tokenIdCall.data?.toString())} 
              name={grantName}
              amount={parseInt(props.grant.claimLimit.toString())} 
              interval={parseInt(props.grant.claimInterval.toString())} 
              closeFn={() => { props.setDialog(null); }} 
            />
          ));
        }}>
          <i className="bi bi-pencil"></i>
        </a>
      </td>
      <td>{toReadableAmount(claimableBalanceCall.data as bigint, chain?.nativeCurrency.decimals, chainConfig.TokenName, 3)}</td>
      <td>{toReadableAmount(totalClaimedCall.data as bigint, chain?.nativeCurrency.decimals, chainConfig.TokenName, 3)}</td>
    </tr>
  )

  
}

export default GrantItem;