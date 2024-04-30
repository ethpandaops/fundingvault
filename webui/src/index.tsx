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

import VaultPage from './components/VaultPage';
import { ephemery } from './config';

const config = getDefaultConfig({
  appName: 'FundingVault',
  projectId: '4b8923523ec77b9be8ab9fd4ff539b48',
  chains: [
    holesky,
    sepolia,
    ephemery,
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
          <VaultPage />
        </RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  </React.StrictMode>
);

reportWebVitals();
