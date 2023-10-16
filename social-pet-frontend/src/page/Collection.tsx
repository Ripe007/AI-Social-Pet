import Collection, { CollectionItem } from "@/components/card/collection";
import Chat from "@/components/card/collection/Chat";
import Food from "@/components/card/collection/food";
import { collection, mintStore } from "@/store/mint.contract";
import { useWallet } from "@/store/wallet";
import { NForm, NFormItem, NGrid, NInput } from "naive-ui";
import { defineComponent, ref } from "vue";
import styled from "vue3-styled-component";

const Collections = () => {
  const Wrapper = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 16px 0;

    h1 {
      font-weight: bold;
    }

    p {
      text-indent: 0.5em;
    }
  `;
  const Flex = styled.div`
    display: flex;
    align-items: center;
    align-self: flex-start;
  `;
  const Avant = styled<any>("img")`
    width: ${(props) => (props.size ? props.size + "px" : "60px")};
    height: ${(props) => (props.size ? props.size + "px" : "60px")};
    border-radius: 50%;
    margin-right: 12px;
    border: 1px solid #e6dbdb;
  `;
  const OverflowImg = styled("img")`
    width: ${(props) => (props?.width ? props.width + "px" : "200px")};
    height: ${(props) => (props?.width ? props.width + "px" : "200px")};
    box-shadow: 0px 0px 10px 0px #e6dbdb;
    margin: 20px auto;
    border-radius: 15px;
  `;
  const OverflowDiv = styled.div`
    width: 200px;
    height: 200px;
    box-shadow: 0px 0px 10px 0px #e6dbdb;
    margin: 20px auto;
    border-radius: 15px;
    display: flex;
    align-items: center;
    justify-centent: center;
  `;
  const ButGroup = styled(Flex)`
    margin: 20px auto;
    width: 100%;
  `;
  const Btn = styled.div`
    width: 100px;
    line-height: 40px;
    display: flex;
    align-items: center;
    align-self: flex-start;
    position: relative;

    img {
      width: 20px;
    }
  `;

  const ActiveBlock = styled.div`
    position: absolute;
    width: 100px;
    height: 5px;
    background: repeating-linear-gradient(
      56deg,
      rgba(68, 206, 246, 0.5),
      rgba(68, 206, 246, 0.5) 14px,
      white 16px,
      white 22px
    );
    bottom: 0;
  `;

  const List = styled.ul`
    width: 100%;
  `;
  const ListItem = styled.li`
    display: flex;
    align-items: center;
    justify-content: center;
  `;

  const FlexRelative = styled.div`
    position: relative;
  `;
  const TalkRoundBtn = styled.div`
    background: rgb(101 219 181);
    border-radius: 5px;
    position: absolute;
    top: 15px;
    right: -40px;
    color: #fff;
    padding: 0 10px;
  `;
  return defineComponent({
    setup() {
      const chating = ref(false);
      const chatRef = ref<any>(null);
      const useMintStore = mintStore();
      const wallet = useWallet();
      const type = ref(0);
      const chageType = async (value: number) => {
        type.value = value;
      };
      const target = ref<collection | null>(null);
      const selectCollection = (item: collection) => {
        target.value = item;
      };
      const open = () => {
        chating.value = true;
        chatRef.value?.pushTarget(target.value);
        chatRef.value?.open();
      };
      return () => (
        <Wrapper>
          <Chat
            ref={chatRef}
            onChange={(e: any) => {
              chating.value = e;
            }}
          />
          <Flex>
            <Avant src="/assets/logo.png"></Avant>
            <div>
              <h1>
                {wallet.wallet.account.substring(0, 6) +
                  "..." +
                  wallet.wallet.account.substring(38)}
              </h1>
              <p>collection:{useMintStore.collection.length}</p>
              <p>balance:{wallet.wallet.balance.div(1e18).toFixed(3)} ETH</p>
            </div>
          </Flex>
          {target.value ? (
            <FlexRelative>
              <TalkRoundBtn
                class="animate__animated animate__pulse animate__fast"
                onClick={open}
              >
                Talk With Pet
              </TalkRoundBtn>
              <OverflowImg src={target.value.image}></OverflowImg>
            </FlexRelative>
          ) : (
            <OverflowDiv>
              <OverflowImg src="/assets/loader.gif" width={100}></OverflowImg>
            </OverflowDiv>
          )}
          {target.value && (
            <List>
              {target.value.attributes.map((attr) => (
                <ListItem>
                  <h4>{attr.TraitType}: </h4>
                  <h5>{attr.value}</h5>
                </ListItem>
              ))}
              <ListItem>
                <h4>name: </h4>
                <h5>{target.value.name}</h5>
              </ListItem>
              <ListItem>
                <h4>description: </h4>
                <h5>
                  {target.value.description
                    ? target.value.description
                    : "not description"}
                </h5>
              </ListItem>
            </List>
          )}

          {/* <ButGroup> */}
          <Btn onClick={chageType.bind(null, 0)}>
            <img src="/assets/icon/vitamins.svg" alt="" />
            Collections
          </Btn>
          <Collection onChange={selectCollection} />

          <Btn onClick={chageType.bind(null, 1)}>
            <img src="/assets/icon/food.svg" />
            Food
          </Btn>
          <Food />
        </Wrapper>
      );
    },
  });
};

export default Collections();
