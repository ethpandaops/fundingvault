import '@rainbow-me/rainbowkit/styles.css';
import './utils/polyfills';
import './index.css';
import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './utils/reportWebVitals';
import { getDefaultConfig, RainbowKitProvider } from '@rainbow-me/rainbowkit';
import { WagmiProvider } from 'wagmi';
import { ConnectButton } from '@rainbow-me/rainbowkit';
import { Suspense } from "react";
import { BrowserRouter } from "react-router-dom";
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import {
  holesky,
  sepolia,
} from 'wagmi/chains';

import Background from "./components/background/Background";
import VaultInfo from './components/vaultinfo/VaultInfo';

const config = getDefaultConfig({
  appName: 'RainbowKit demo',
  projectId: 'YOUR_PROJECT_ID',
  chains: [
    holesky,
    sepolia,
  ],
});

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const queryClient = new QueryClient();

root.render(
  <React.StrictMode>
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <RainbowKitProvider>
          <BrowserRouter>
            <Suspense fallback={null}>
              <Background>
                <div className="foreground-container">
                  <div className="page-header">
                    <ConnectButton />
                  </div>
                  <div className="page-wrapper container-fluid">
                    <VaultInfo />
                  </div>
                  <div className="page-footer">
                    <span>Powered by <a href="https://github.com/pk910/holesky-fundingvault" target="_blank">pk910/holesky-fundingvault</a> | git-xxxxxx</span>
                  </div>
                </div>
              </Background>
            </Suspense>
          </BrowserRouter>
        </RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  </React.StrictMode>
);

reportWebVitals();
