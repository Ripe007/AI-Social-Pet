import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
require('dotenv').config();

const PRIVATE_KEY = process.env.PRIVATE_KEY;
const ENDPOINT_URL = process.env.ENDPOINT_URL;

const config: HardhatUserConfig = {
  defaultNetwork: "goerli",
  solidity: {
    version: "0.8.20",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    goerli: {
      url: ENDPOINT_URL,
      accounts: [`0x${PRIVATE_KEY}`],
      chainId: 5,
    },
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_APIKEY,
  },
  
};

export default config;
