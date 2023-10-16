import { defineStore } from "pinia";
import { useWallet } from "./wallet";
import { ethers } from "ethers";
import Erc1155 from "@/utils/abi/Erc1155.json";
import BigNumber from "bignumber.js";

export const useTip = defineStore("tip", () => {
  const tip = async (
    foodAmount: number | string,
    ethAmount: number | string,
    to: string
  ) => {
    const wallet = useWallet();
    const [provider, singer, account] = await wallet.connect();
    const FoodProvider = new ethers.Contract(
      (import.meta as any).env.VITE_FOOD,
      Erc1155,
      singer
    );
    if (foodAmount) {
      const cost = await FoodProvider.costById(1);
      const bool = await FoodProvider.mint(to, 1, foodAmount, {
        value: cost.toString(),
      });
      await bool.wait();
    }
    if (ethAmount) {
      const tx = {
        from: account,
        to,
        value: BigNumber(ethAmount).times(1e18).toFixed(0),
      };
      const bool = await singer.sendTransaction(tx);
      await bool.wait();
    }
  };
  return { tip };
});
