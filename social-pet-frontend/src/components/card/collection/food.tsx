import { defineComponent, ref } from "vue";
import styled from "vue3-styled-component";

const Food = () => {
  const Grid = styled.div`
    display: flex;
    overflow: auto;

    &::-webkit-scrollbar {
        display: none;
      }
  `;
  const Flex = styled.div`
    img {
      width: 138px;
      height: 138px;
    }
  `;
  const Tokenid = styled.div`
    position: absolute;
    bottom: 10px;
    right: 10px;
  `;
  return defineComponent({
    setup() {
      const items = ref([
        { img: "/assets/food.png", tokenId: 1 },
        { img: "/assets/food.png", tokenId: 2 },
        { img: "/assets/food.png", tokenId: 3 },
        { img: "/assets/food.png", tokenId: 4 },
      ]);

      return () => (
        <Grid>
          {items.value.map((item) => (
            // <Flex>
            <img src={item.img} width={130} />
            // <Tokenid>#{item.tokenId}</Tokenid>
            // </Flex>
          ))}
        </Grid>
      );
    },
  });
};
export default Food();
