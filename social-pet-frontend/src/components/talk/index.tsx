import { useTip } from "@/store/tip.contract";
import {
  NCard,
  NDropdown,
  NGi,
  NGrid,
  NInputNumber,
  NModal,
  NSpin,
  useMessage,
} from "naive-ui";
import { defineComponent, ref } from "vue";
import { useRouter } from "vue-router";
import styled from "vue3-styled-component";

export type TalkSource = {
  userAddr: string;
  context: string;
  images: string[];
  source_id: string;
  updated_at: string;
  like_count: number;
  comment_count: number;
  forward_count: number;
};
const Talk = () => {
  const FlexCol = styled.div`
    display: flex;
    padding: 12px;
    margin-bottom: 20px;

    &:hover {
      background: #ededed;
    }
  `;
  const Avanent = styled.img`
    width: 60px;
    height: 60px;
    border-radius: 50%;
  `;
  const Content = styled.div`
    display: flex;
    flex-direction: column;
    text-align: left;
    flex: 1;
  `;
  const NickName = styled.h2`
    font-size: 16px;
    font-weight: bold;
  `;
  const Text = styled.span``;
  const ImgGroup = styled<any>("div")`
    display: grid;
    grid-template-columns: repeat(
      ${(props) => (props.len > 3 ? 3 : props.len)},
      1fr
    );
    grid-gap: 5px 5px;
  `;
  const Flex = styled.div`
    display: flex;
    justify-content: space-between;
    margin-top: 5px;
  `;
  const But = styled.div`
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    img {
      margin-right: 5px;
    }
  `;
  const Card = styled.div`
    width: 80vw;
    margin: 0 auto;
    background: #fff;
    display: flex;
    flex-direction: column;
    border-radius: 10px;
    padding: 12px;
    position: relative;
  `;
  const closeIcon = styled.img`
    postion: absolute;
    width: 30px;
    height: 30px;
    top: 12px;
    right: 12px;
  `;
  const Bold = styled.span`
    font-size: 30px;
    color: #000;
    font-family: serif;
    font-weight: bold;
  `;
  const H2Span = styled(Bold)`
    font-size: 24px;
    margin-top: 20px;
    margin-right: 10px;
  `;
  const GiFlexCol = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-centent: center;
  `;
  const TipBtn = styled.div`
    background: rgb(101 219 181);
    color: #fff;
    padding: 3px 15px;
    border-radius: 5px;
    width: 100px;
    text-align: center;
    margin: 16px auto;
  `;

  return defineComponent({
    props: [
      "userAddr",
      "context",
      "images",
      "source_id",
      "updated_at",
      "like_count",
      "comment_count",
      "forward_count",
    ],
    setup(props, { expose }) {
      const router = useRouter();
      let source = props as TalkSource;
      const showDropdown = ref(false);
      const showModal = ref(false);
      const loading = ref(false);
      const ethRef = ref<any>(null);
      const foodRef = ref<any>(null);
      const tipStore = useTip();
      const message = useMessage();
      const tipSource = ref({
        food: "",
        eth: "",
        to: "",
      });
      const Options = ref([
        {
          icon: () => <img src="/assets/food.png" width={35}></img>,
          label: "Food",
          key: "food",
        },
        {
          icon: () => <img src="/assets/icon/eth.webp" width={35}></img>,
          label: "ETH",
          key: "ETH",
        },
      ]);
      const handleClick = () => {
        showDropdown.value = !showDropdown.value;
      };
      const handleSelect = (item: any) => {
        showModal.value = true;
        tipSource.value.food = "";
        tipSource.value.eth = "";
        // setTimeout(() => {
        //   if (e === "food") foodRef.value?.focus();
        //   else ethRef.value?.focus();
        // }, 1000);
        // console.log(item);
      };
      const goto = (item: any) => {
        router.push({
          path: "/detail",
          query: {
            id: source.source_id,
          },
        });
      };
      const comfirm = () => {
        if (tipSource.value.food || tipSource.value.eth) {
          loading.value = true;
          tipStore
            .tip(tipSource.value.food, tipSource.value.eth, source.userAddr)
            .then((res) => {
              console.log("res: ", res);
              message.success("Successful!");
              loading.value = false;
            })
            .catch((err) => {
              console.log("err: ", err);
              message.error("Faild!");
              loading.value = false;
            });
        }
      };
      return () => (
        <FlexCol>
          <Avanent src="/assets/logo.png"></Avanent>
          <Content onClick={goto}>
            <NickName>
              {source.userAddr.substring(0, 6) +
                "..." +
                source.userAddr.substring(38)}
            </NickName>
            <Text>{source.context}</Text>
            <ImgGroup len={source.images.length}>
              {source.images.map((item) => (
                <img src={item} key={item}></img>
              ))}
            </ImgGroup>
            <Flex>
              <But>
                <img src="/assets/icon/msg.svg" width={20} alt="" />
                <div>{source.comment_count}</div>
              </But>
              <But>
                <img src="/assets/icon/back.svg" width={20} alt="" />
                <div>{source.forward_count}</div>
              </But>
              <But
                onClick={(e) => {
                  e.stopPropagation();
                }}
              >
                <img src="/assets/icon/donate.svg" width={20} alt="" />
                <NDropdown
                  show={showDropdown.value}
                  options={Options.value}
                  onSelect={handleSelect}
                >
                  <div onClick={handleClick}>tip</div>
                </NDropdown>
              </But>
            </Flex>
          </Content>

          <NModal v-model:show={showModal.value} auto-focus={false}>
            <Card>
              <H2Span>Tip!</H2Span>
              <NGrid cols={2} xGap={24}>
                <NGi>
                  <GiFlexCol>
                    <img src="/assets/food.png" width={80} alt="" />
                    <NInputNumber
                      ref={foodRef.value}
                      v-model:value={tipSource.value.food}
                      placeholder=""
                      format={(value) =>
                        value ? parseInt(value.toString()).toString() : ""
                      }
                      clearable
                    />
                  </GiFlexCol>
                </NGi>

                <NGi>
                  <GiFlexCol>
                    <img src="/assets/icon/eth.webp" width={80} alt="" />
                    <NInputNumber
                      v-model:value={tipSource.value.eth}
                      ref={ethRef.value}
                      placeholder=""
                      clearable
                      step={0.05}
                    />
                  </GiFlexCol>
                </NGi>
              </NGrid>
              <TipBtn onClick={comfirm}>
                <NSpin show={loading.value} size="small">
                  Confirm
                </NSpin>
              </TipBtn>
            </Card>
          </NModal>
        </FlexCol>
      );
    },
  });
};

export default Talk() as any;
