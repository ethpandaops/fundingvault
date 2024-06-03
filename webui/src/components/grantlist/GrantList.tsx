import {
	useAccount,
  useReadContract,
} from "wagmi";
import { ConfigForChainId } from "../../utils/chaincfg";

import FundingVaultAbi from "../../abi/FundingVault.json";
import CurrentConfig from "../../config";
import { useEffect, useState } from "react";
import { toReadableDuration } from "../../utils/ConvertHelpers";

import "./GrantList.css"
import GrantItem from "./GrantItem";
import GrantCreate from "../grant_create/GrantCreate";


const GrantList = (): React.ReactElement => {
  const { address, chain } = useAccount();
  let chainConfig = ConfigForChainId(chain!.id)!;
  let [managerDialog, setManagerDialog] = useState<React.ReactElement>(null);
  
  const managerCooldownSettings = useReadContract({
		address: chainConfig.VaultContractAddr,
    account: address,
		abi: FundingVaultAbi,
		chainId: chainConfig.Chain.id,
		functionName: "getManagerGrantLimits",
		args: [ ],
	});
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
  const adminCheck = useReadContract(chain ? {
    address: chainConfig.VaultContractAddr,
    account: address,
    abi: FundingVaultAbi,
    chainId: chainConfig.Chain.id,
    functionName: "hasRole",
    args: [ CurrentConfig.AdminRole, address ],
  }: undefined);

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
  }, [managerCooldown, grantList]);

  //console.log(grantList.data)

  var grantListEls: React.ReactElement[] = [];
  if(Array.isArray(grantList.data)) {
    grantList.data.forEach((grant, index) => {
      grantListEls.push(<GrantItem tokenIdx={index} grant={grant} setDialog={setDialog} />)
    });
  }

  var managerCooldownData = {
    locked: false,
    admin: false,
    cooldown: 0,
    threshold: 0,
    percent: 0,
    quotaAmount: 0,
    quotaTime: 0,
    lockQuota: 0,
  };
  if(adminCheck.isFetched && managerCooldown.isFetched && managerCooldownSettings.isFetched) {
    let cooldownData = managerCooldown.data as bigint;
    let cooldownSettings = managerCooldownSettings.data as [bigint, bigint, number, number];

    managerCooldownData.admin = adminCheck.data as boolean;
    managerCooldownData.cooldown = parseInt(cooldownData.toString());
    managerCooldownData.quotaAmount = parseInt(cooldownSettings[0].toString());
    managerCooldownData.quotaTime = parseInt(cooldownSettings[1].toString());
    managerCooldownData.lockQuota = cooldownSettings[2];
    managerCooldownData.threshold = cooldownSettings[3];
    managerCooldownData.locked = managerCooldownData.cooldown > managerCooldownData.threshold;
    managerCooldownData.percent = managerCooldownData.locked ? 100 : Math.floor(100 / managerCooldownData.threshold * managerCooldownData.cooldown);
    console.log(managerCooldownData)
  }

  return (
    <div className="page-block grants-page">
      <h1>Grants List</h1>
      <div className="manager-limits">
        Manager Cooldown:
        {managerCooldownData.admin ? <span className="mx-2 badge text-bg-success">Admin</span> : null}

        <div className="progress cooldown-bar">
          <div 
            className={["progress-bar", "cooldown-bar-value", managerCooldownData.locked ? "bg-warning" : ""].join(" ")} 
            role="progressbar" style={{"width": managerCooldownData.percent + "%"}} 
            aria-valuenow={managerCooldownData.percent} aria-valuemin={0} aria-valuemax={100}
          >
            {managerCooldownData.locked ? 
            <div>
              locked for {toReadableDuration(managerCooldownData.cooldown - managerCooldownData.threshold)}
            </div> :
            <div>
              {managerCooldownData.percent} % ({toReadableDuration(managerCooldownData.cooldown)})
            </div> }
          </div>
        </div>
      </div>

      <div className="grants-list mt-2">
        <div className="create-btn">
          <button className="btn btn-primary" onClick={() => {
            setDialog((
              <GrantCreate closeFn={() => { setDialog(null); }} />
            ));
          }}>New Grant</button>
        </div>
        <div className="pt-3">All existing grants:</div>
        <table className="table grants-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Owner</th>
              <th>Name</th>
              <th>Granted Amount</th>
              <th>Claimable</th>
              <th>Total Claimed</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {grantList.isFetched ? grantListEls : <tr><td colSpan={7}>Loading...</td></tr>}
          </tbody>
        </table>
      </div>
      
      {managerDialog}
    </div>
  )

  function setDialog(element: React.ReactElement) {
    setManagerDialog(element);
  }
  
}

export default GrantList;