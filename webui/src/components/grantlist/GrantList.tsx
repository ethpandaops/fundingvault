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

import "./GrantList.css"
import GrantItem from "./GrantItem";
import GrantCreate from "../grant_create/GrantCreate";


const GrantList = (): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [managerDialog, setManagerDialog] = useState<React.ReactElement>(null);
  
  const managerCooldown = useReadContract({
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getManagerCooldown",
		args: [ address ],
	});
  const grantList = useReadContract({
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getGrants",
		args: [ ],
	});

  //console.log(grantDetails.data);
  useEffect(() => {
    const interval = setInterval(() => {
      console.log("refetch");
      managerCooldown.refetch();
      grantList.refetch();
    }, 15000);
    return () => {
      clearInterval(interval);
    };
  }, []);

  console.log(grantList.data)

  var grantListEls: React.ReactElement[] = [];
  if(Array.isArray(grantList.data)) {
    grantList.data.forEach((grant, index) => {
      grantListEls.push(<GrantItem tokenIdx={index} grant={grant} setDialog={setDialog} />)
    });
  }

  return (
    <div className="page-block grants-page">
      <h1>Grants List</h1>
      <div className="create-btn">
        <button className="btn btn-primary" onClick={() => {
          setDialog((
            <GrantCreate closeFn={() => { setDialog(null); }} />
          ));
        }}>New Grant</button>
      </div>
      <p>
        All existing grants:
      </p>
      
      <table className="table grants-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Owner</th>
            <th>Name</th>
            <th>Granted Amount</th>
            <th>Claimable</th>
            <th>Total Claimed</th>
          </tr>
        </thead>
        <tbody>
          {grantList.isFetched ? grantListEls : <tr><td colSpan={4}>Loading...</td></tr>}
        </tbody>
      </table>
      
      {managerDialog}
    </div>
  )

  function setDialog(element: React.ReactElement) {
    setManagerDialog(element);
  }
  
}

export default GrantList;