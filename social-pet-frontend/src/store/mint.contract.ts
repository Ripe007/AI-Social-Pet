import { defineStore } from "pinia";
import { useWallet } from "./wallet";
import MintNftAbi from "@/utils/abi/MintNFT.json";
import { ethers } from "ethers";
import BigNumber from "bignumber.js";
import { ref } from "vue";
import axios from "axios";

type attributes = {
  value: string;
  TraitType: string;
};
export type collection = {
  tokenid: any;
  attributes: attributes[];
  image: string;
  name: string;
  description: string;
};
export const mintStore = defineStore("mint", () => {
  let timer: NodeJS.Timeout | null;
  const collection = ref<collection[]>([]);

  const mint = async (uri: string) => {
    const wallet = useWallet();
    const [provider, singer, account] = await wallet.connect();
    const MintNftProvider = new ethers.Contract(
      (import.meta as any).env.VITE_MINE_SOL,
      MintNftAbi,
      singer
    );
    const price = await MintNftProvider.cost();

    return MintNftProvider.mint(
      account,
      (import.meta as any).env.VITE_IMG_URI + uri,
      {
        value: price.toString(),
      }
    );
  };

  const NextTokenId = async () => {
    const wallet = useWallet();
    const [provider, singer, account] = await wallet.connect();
    const MintNftProvider = new ethers.Contract(
      (import.meta as any).env.VITE_MINE_SOL,
      MintNftAbi,
      provider
    );
    return (await MintNftProvider.getNextId()).toNumber() + 1;
  };
  const init = async () => {
    const wallet = useWallet();
    const [provider, singer, account] = await wallet.connect();
    const MintNftProvider = new ethers.Contract(
      (import.meta as any).env.VITE_MINE_SOL,
      MintNftAbi,
      provider
    );
    async function func() {
      const tokenIds: any[] = await MintNftProvider.tokensOfOwner(account);
      if (tokenIds.length) {
        tokenIds.forEach(async (tokenid) => {
          let index = collection.value.findIndex(
            (item) => item.tokenid === tokenid.toNumber()
          );
          if (index === -1) {
            try {
              let json: string = await MintNftProvider.tokenURI(
                tokenid.toNumber()
              );
              json = json.replace("/static", "");
              json = json.replace(".comc", ".com");
              const resource = await axios.get(json);
              collection.value.push({
                tokenid: tokenid.toNumber(),
                image: resource.data.image,
                name: resource.data.name,
                description: resource.data.description,
                attributes: resource.data.attributes as attributes[],
              });
            } catch (err) {}
          }
        });
      }
    }
    func();
    if (timer) clearInterval(timer);
    timer = setInterval(func, 10 * 1000);
  };
  return { mint, NextTokenId, init, collection };
});
