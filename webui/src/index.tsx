import '@rainbow-me/rainbowkit/styles.css';
import './utils/polyfills';
import './index.css';
import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './utils/reportWebVitals';
import { getDefaultConfig, RainbowKitProvider } from '@rainbow-me/rainbowkit';
import { WagmiProvider } from 'wagmi';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import VaultPage from './components/VaultPage';
import { KnownChains } from './config';

const config = getDefaultConfig({
  appName: 'FundingVault',
  projectId: 'eacf22cd196bfe746ffa0f7e9c275339',
  chains: KnownChains as any,
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
