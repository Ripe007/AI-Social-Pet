import config from "../../setting.json";
import erc20Abi from "@/utils/abi/erc20.json";
import routerAbi from "@/utils/abi/ROUTER_PANCAKE.json";
import swapAbi from "@/utils/abi/SWAP_CONTRACT.json";
import { ethers } from "ethers";
import BigNumber from "bignumber.js";
import moment from "moment";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useWallet = defineStore("mywallet", () => {
  const wallet = ref({
    account: "0x0000000000000000000000000000000000000000",
    balance: BigNumber(0),
    isConnect: false,
  });

  const getBalanceOf = async (
    provider: ethers.providers.Web3Provider,
    account: string = wallet.value.account
  ) => {
    if (account === wallet.value.account) {
      const value = await provider.getBalance(account);
      wallet.value.balance = BigNumber(value.toString());
    } else {
      const value = await provider.getBalance(account);
      return BigNumber(value.toString());
    }
  };

  const connect = async (): Promise<
    [ethers.providers.Web3Provider, ethers.providers.JsonRpcSigner, string]
  > => {
    await (window as any).ethereum.request({
      method: "eth_requestAccounts",
    });
    const provider = new ethers.providers.Web3Provider(
      (window as any).ethereum
    );
    const singer = provider.getSigner();
    const account = await singer.getAddress();
    wallet.value.account = account;
    wallet.value.isConnect = true;
    getBalanceOf(provider);
    return [provider, singer, account];
  };

  const sign = async () => {
    await (window as any).ethereum.request({
      method: "personal_sign",
      params: [wallet.value.account, "sign transition"],
    });
    return true;
  };
  return { wallet, connect, sign, getBalanceOf };
});
