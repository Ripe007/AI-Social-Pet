import { collection } from "@/store/mint.contract";
import axios from "axios";
import { NDrawer, NDrawerContent, NInput, NModal } from "naive-ui";
import { VNode, VueElement, computed, defineComponent, ref, watch } from "vue";
import styled from "vue3-styled-component";

type talker = {
  type: string; // ai,man
  value: string | VNode;
  img: string;
  name: string;
};
const host = "http://AISound.aipet.vip:8000";

const Chat = () => {
  const Col = styled.div`
    display: flex;
    flex-direction: column;
    max-height: 265px;
    overflow-y: auto;
  `;
  const TalkCard = styled<any>("div")`
    clear: both;
    display: flex;
    justify-content: ${(props) => props?.placement ?? "flex-start"};

    img {
      width: 50px;
      height: 50px;
      margin-top: 5px;
      border-radius: 50%;
    }
    .group {
      display: column;
      margin-left: 10px;
    }
    .name {
    }
    .context {
      background: rgb(101 219 181);
      color: #fff;
      padding: 0 8px;
      border-radius: 4px;
      max-width: 150px;
      min-height: 2em;
      min-width: 30%;
    }
  `;
  const FixedBottom = styled.div`
    position: fixed;
    bottom: 20px;
    left: 16px;
    right: 16px;
  `;

  return defineComponent({
    setup(props: any, { expose, emit }) {
      const show = ref(false);
      const target = ref<any>({});
      const context = ref<talker[]>([]);
      const sendMessage = ref(null);
      const aiLoading = ref(false);

      const Send = () => {
        if (sendMessage.value) {
          context.value.push({
            type: "man",
            value: sendMessage.value ?? "",
            img: "/assets/logo.png",
            name: "",
          });
          let value = sendMessage.value;
          sendMessage.value = null;

          let index: number;
          setTimeout(() => {
            index = context.value.length;
            context.value.push({
              type: "ai",
              value: <img src="/assets/loader.gif" width={40} />,
              img: target.value.image,
              name: target.value.attributes[0].value,
            });
          }, 1000);
          axios
            .get(host + `/api/response?message=${value}`, {
              headers: {
                token: "1c98aa19e8e27ba0a49c5dca57d16d2d",
              },
            })
            .then((res) => {
              let resp = res.data;
              let msg = "";
              if (resp.status === "ok") {
                msg = resp.data.msg;
                context.value[index].value = msg;
              }
            });
        }
      };
      watch(show, () => {
        console.log(target.value);
        if (show.value === false) {
          context.value = [];
        }
        emit("change", show.value);
      });
      expose({
        open: () => {
          show.value = true;
        },
        pushTarget: (iTarget: any) => {
          target.value = iTarget;
          context.value.push({
            type: "ai",
            value: "Master, I am your little AI pet, come and play with me",
            img: target.value.image,
            name: target.value.attributes[0].value,
          });
        },
      });
      return () => (
        <NDrawer v-model:show={show.value} placement="bottom" height={350}>
          <NDrawerContent>
            <Col>
              {context.value.map((talker) => (
                <TalkCard
                  placement={talker.type === "ai" ? "flex-start" : "flex-end"}
                  style={{ marginTop: talker.type !== "ai" ? "10px" : "0" }}
                >
                  <img
                    src={talker.img}
                    alt=""
                    style={{ order: talker.type === "ai" ? 1 : 2 }}
                  />
                  <div
                    class="group"
                    style={{ order: talker.type === "ai" ? 2 : 1 }}
                  >
                    <div class="name">{talker.name}</div>
                    <div class="context">{talker.value}</div>
                  </div>
                </TalkCard>
              ))}
            </Col>
            <FixedBottom>
              <NInput
                v-model:value={sendMessage.value}
                type="text"
                placeholder=""
                v-slots={{
                  suffix: () => <div onClick={Send}>Send</div>,
                }}
              ></NInput>
            </FixedBottom>
          </NDrawerContent>
        </NDrawer>
      );
    },
  });
};

export default Chat() as any;
