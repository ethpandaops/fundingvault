import React from "react";
import { useAccount, useReadContract } from 'wagmi';
import { ConnectButton } from '@rainbow-me/rainbowkit';
import { Suspense } from "react";
import { BrowserRouter, Route, Routes, Link } from "react-router-dom";

import CurrentConfig from "../config";
import { ConfigForChainId } from "../utils/chaincfg";
import FundingVaultAbi from "../abi/FundingVault.json";

import Background from "./background/Background";
import VaultInfo from './vaultinfo/VaultInfo';
import GrantList from "./grantlist/GrantList";

const VaultPage = (): React.ReactElement => {
  const { address: walletAddress, isConnected, chain } = useAccount();

  let chainConfig = ConfigForChainId(chain?.id)!;
  const managerCheck = useReadContract(chain ? {
    address: chainConfig.VaultContractAddr,
    account: walletAddress,
    abi: FundingVaultAbi,
    chainId: chainConfig.Chain.id,
    functionName: "hasRole",
    args: [ CurrentConfig.ManagerRole, walletAddress ],
  }: undefined);

  var isManager = isConnected && chain && managerCheck.isSuccess && managerCheck.data;

  return (
    <BrowserRouter>
      <Suspense fallback={null}>
        <Background>
          <div className="foreground-container">
            <div className="page-header">
              {isManager ? 
                <div className="manager-btn px-3">
                  <Link to="/manage">
                    <button type="button" className="btn btn-light font-weight-bold">Manage Grants</button>
                  </Link>
                </div>
              : null}
              <ConnectButton />
            </div>
            <div className="page-wrapper container-fluid">
              <Routes>
                <Route path="/" element={(
                  <VaultInfo />
                )} />
                <Route path="/manage" element={isManager ? (
                  <GrantList />
                ) : null} />
              </Routes>
            </div>
            <div className="page-footer">
              <span>Powered by <a href="https://github.com/pk910/holesky-fundingvault" target="_blank" rel="noreferrer">pk910/holesky-fundingvault</a> | {CurrentConfig.AppVersion ? "git-" + CurrentConfig.AppVersion : "dev build"}</span>
            </div>
          </div>
        </Background>
      </Suspense>
    </BrowserRouter>
  )
}

export default VaultPage;