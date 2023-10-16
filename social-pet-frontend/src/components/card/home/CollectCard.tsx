import {
  NForm,
  NFormItemGi,
  NGi,
  NGrid,
  NIcon,
  NImage,
  NInput,
  NModal,
  NUpload,
  useMessage,
} from "naive-ui";
import { defineComponent, inject, ref } from "vue";
import styled from "vue3-styled-component";
import axios from "axios";
import { TrashBinOutline } from "@vicons/ionicons5";
import { FormInst, FormItemRowRef } from "naive-ui/es/form/src/interface";
import { mintStore } from "@/store/mint.contract";
import { useWallet } from "@/store/wallet";
import { useRouter } from "vue-router";

const CollectCard = () => {
  const FlexCol = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  `;
  const Card = styled(FlexCol)`
    width: 100%;
    min-height: 200px;
    box-shadow: 0px 0px 10px 0px #e6dbdb;
    border-radius: 15px;
    background: #fff;
    padding: 16px 0;
  `;
  const ImgCard = styled(Card)`
    width: 90%;
    position: relative;
    padding: 45px 25px 16px;
    display: flex;
    justify-content: center;
  `;
  const PositionImg = styled.img`
    position: absolute;
    z-index: 1;
  `;
  const CollectImage = styled(NImage)`
    width: 250px;
    border-radius: 15px;
    box-shadow: 0px 0px 5px 0px #000;
  `;
  const CloseIcon = styled.img`
    width: 25px;
    height: 25px;
    position: absolute;
    top: 12px;
    right: 12px;
  `;
  const Round = styled(FlexCol)`
    width: 120px;
    height: 120px;
    border-radius: 50%;
    border: 1px solid #e6dbdb;
    padding: 10px;
    box-shadow: 0px 0px 10px 0px #e6dbdb;
  `;
  const OpacityImg = styled.img`
    filter: opacity(0.5);
  `;
  const List = styled.div`
    display: flex;
    flex-direction: column;
  `;

  const Item = styled.div`
    display: flex;
    justify-content: space-between;
    padding: 5px 16px;

    div {
      font-weight: 500;
    }
  `;

  const Icon = styled.img`
    width: 18px;
    height: 18px;
    margin-right: 5px;
  `;

  const Upload = styled(NUpload)`
    display: flex;
    justify-content: center;
  `;

  const But = styled.div`
    background: rgb(101 219 181);
    color: #fff;
    padding: 3px 15px;
    border-radius: 5px;
  `;

  return defineComponent({
    setup() {
      const pageLoader = inject("pageLoader", (status: boolean) => {});
      const MintStore = mintStore();
      const wallet = useWallet();
      const message = useMessage();
      const router = useRouter();
      const show = ref(false);
      const loading = ref(false);
      const uploadUrl = ref(null);

      // const show = ref(true);
      // const loading = ref(false);
      // const uploadUrl = ref("/static/draw/1713601309811321179.png");

      const petInfo = ref({
        nickname: "",
        description: "",
        identify: "", // 物种类型
      });
      const formRef = ref<FormInst | null>(null);
      /** ai生成图片 */
      const customRequest = ({ file }: any) => {
        if (loading.value) return;
        show.value = true;
        loading.value = true;
        const reader = new FileReader();

        /** 轮询 */
        function getTask(task_id: any) {
          return new Promise(async (resolve, reject) => {
            while (true) {
              try {
                const { data } = await axios.post("/api/draw/get", {
                  task_id,
                });
                if (data.data.status === 1) {
                  resolve(data.data);
                  break;
                }
              } catch (err) {
                console.log("err :", err);
                continue;
              }
            }
          });
        }
        reader.addEventListener("load", async () => {
          const base64 = reader.result;
          while (true) {
            try {
              const { data } = await axios.post("/api/draw/add", {
                prompt: "宠物",
                image: base64,
              });
              {
                if (data.code === 500) continue;
              }
              const task_id = data.data.task_id;
              /** 固定task——id */
              // const task_id = "1713601309811321179";
              const { image_url, animal_identify } = (await getTask(
                task_id
              )) as any;
              uploadUrl.value = image_url;
              petInfo.value.identify = animal_identify;
              break;
            } catch (err) {
              continue;
            }
          }
        });
        reader.readAsDataURL(file.file);
      };
      /** 铸币 */
      const mint = async () => {
        // const data = axios.get('https://aipet.hm-swap.com/json/9.json')
        // return;
        if (!petInfo.value.nickname) return;
        pageLoader(true);
        try {
          const nextTokenId = await MintStore.NextTokenId();
          const { data } = await axios.post("/api/upload/json", {
            metadata: {
              attributes: [
                {
                  value: petInfo.value.nickname,
                  trait_type: "nickname",
                },
              ],
              description: petInfo.value.description,
              image: (import.meta as any).env.VITE_IMG_URI + uploadUrl.value,
              name: "OG Social Pet #" + nextTokenId,
            },
            token_id: nextTokenId,
          });
          if (data.code === 0) {
            const tx = await MintStore.mint(data.data);
            await tx.wait();
            await axios.post("/api/nft/mint", {
              tx_id: tx.hash,
              user_addr: wallet.wallet.account,
              identify: "金毛",
            });
            message.success("Mint!");
            setTimeout(() => {
              show.value = false;
              pageLoader(false);
              router.push("/collection");
            }, 1500);
          }
        } catch (err) {
          console.log("error :", err);
        }
      };
      return () => (
        <Card>
          <Upload show-file-list={false} customRequest={customRequest}>
            <Round>
              <OpacityImg src="/assets/banner2.png" alt="" />
              <img src="/assets/icon/upload.svg" width={40} alt="" />
            </Round>
          </Upload>
          <List>
            <Item>
              <Icon src="/assets/icon/item1.svg"></Icon>
              <div>Upload your pet's picture, and craft a unique NFT</div>
            </Item>
            <Item>
              <Icon src="/assets/icon/item2.svg" />
              <div>
                Buy, sell pets, gear, and valuables in the virtual market.
              </div>
            </Item>
            <Item>
              <Icon src="/assets/icon/item3.svg" />
              <div>
                20% of the minted coins will be allocated to the foundation,
                granting you valuable voting rights. Let's use our voting power
                together to help rescue these adorable little ones.
              </div>
            </Item>
            <Item>
              <Icon src="/assets/icon/item4.svg" />
              <div>Real talks with pets; they respond, learn over time</div>
            </Item>
            <Item>
              <Icon src="/assets/icon/item5.svg" />
              <div>
                Connect with others, share pet-raising experiences, and earn
                CROPTY rewards
              </div>
            </Item>
          </List>

          <NModal show={show.value} transformOrigin="center">
            <ImgCard>
              <CloseIcon
                src="/assets/icon/close.svg"
                onClick={() => {
                  show.value = false;
                }}
              />
              {loading.value && (
                <PositionImg src="/assets/loader.gif" width={100} alt="" />
              )}
              {uploadUrl.value && (
                <CollectImage
                  src={uploadUrl.value ?? ""}
                  on-load={() => {
                    loading.value = false;
                  }}
                />
              )}

              {!loading.value && (
                <FlexCol style={{ marginTop: "12px" }}>
                  <NForm
                    labelPlacement="top"
                    model={petInfo.value}
                    ref={formRef}
                  >
                    <NGrid cols={24}>
                      <NFormItemGi span={24} label="Nickname" required>
                        <NInput
                          placeholder="Give your pet a name"
                          v-model:value={petInfo.value.nickname}
                        ></NInput>
                      </NFormItemGi>
                      <NFormItemGi span={25} label="description">
                        <NInput
                          v-model:value={petInfo.value.description}
                          placeholder="Introduce your pet to everyone"
                          type="textarea"
                          round
                          clearable
                          v-slots={{
                            "clear-icon": () => (
                              <NIcon>
                                <TrashBinOutline />
                              </NIcon>
                            ),
                          }}
                        ></NInput>
                      </NFormItemGi>
                    </NGrid>
                  </NForm>
                  <But onClick={mint}>Mint</But>
                </FlexCol>
              )}
            </ImgCard>
          </NModal>
        </Card>
      );
    },
  });
};

export default CollectCard();
