import { mintStore } from "@/store/mint.contract";
import { computed, defineComponent, onMounted, ref } from "vue";
import styled from "vue3-styled-component";
import type { collection } from "@/store/mint.contract";

export type CollectionItem = {
  img: string;
  tokenId: number;
};
const Card = () => {
  const Grid = styled.div`
    display: flex;
    overflow: auto;

    &::-webkit-scrollbar {
      display: none;
    }
  `;
  const Flex = styled.div`
    position: relative;
  `;
  const Img = styled.img`
    width: 138px;
    height: 138px;
    margin: 5px;
    border-radius: 15px;
  `;
  const Tokenid = styled.div`
    position: absolute;
    bottom: 10px;
    right: 10px;
  `;
  return defineComponent({
    setup(props, { emit }) {
      const useMintStore = mintStore();
      const items = computed<collection[]>(() => useMintStore.collection);
      const click = (item: collection) => {
        emit("change", item);
      };
      onMounted(() => {
        emit("change", items.value[0]);
      });
      return () => (
        <Grid>
          {items.value.map((item) => (
              <Img
                src={item.image}
                width={130}
                onClick={click.bind(null, item)}
              />
              // <Tokenid>{item.name}</Tokenid>
          ))}
        </Grid>
      );
    },
  });
};
export default Card() as any;
