import { useWallet } from "@/store/wallet";
import axios from "axios";
import {
  NInput,
  NModal,
  NUpload,
  UploadCustomRequestOptions,
  UploadFileInfo,
} from "naive-ui";
import { defineComponent, inject, ref } from "vue";
import styled from "vue3-styled-component";

const Write = () => {
  const OverCard = styled.div`
    width: 100vw;
    height: 100vh;
    background: #fff;
    padding: 16px;
  `;
  const WrapperHeader = styled.div`
    height: 2em;
    display: flex;
    justify-content: space-between;
    padding: 10px;
  `;
  const WrapperBody = styled.div`
    display: flex;
    flex-direction: column;
    margin: 16px 0;

    .n-input__state-border,
    .n-input__border {
      border: 0 !important;
      box-shadow: none !important;
    }
  `;
  return defineComponent({
    setup(props, ctx) {
      const wallet = useWallet();
      const pageLoader = inject("pageLoader", (status: boolean) => {});
      const show = ref(false);
      const showModalRef = ref(false);
      const previewImageUrlRef = ref("");

      const contextInfo = ref<any>({
        img_urls: [],
        context: "",
        user_addr: "",
      });

      const customRequest = ({
        file,
        data,
        action,
        onFinish,
        onError,
      }: UploadCustomRequestOptions) => {
        const formData = new FormData();
        if (data) {
          Object.keys(data).forEach((key) => {
            formData.append(
              key,
              data[key as keyof UploadCustomRequestOptions["data"]]
            );
          });
        }
        formData.append("file", file.file as File);
        axios
          .post(action as any, formData, {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          })
          .then(({ data }) => {
            contextInfo.value.img_urls.push(data.data);
            onFinish();
          })
          .catch((error) => {
            onError();
          });
      };
      const Done = async () => {
        if (!contextInfo.value.context) {
          return;
        }
        try {
          pageLoader(true);
          contextInfo.value.user_addr = wallet.wallet.account;
          const { data } = await axios.post(
            "/api/community/publish",
            contextInfo.value
          );
          if (data.code === 0) {
            show.value = false;
            pageLoader(false);
          }
        } catch (err) {}
        pageLoader(false);
      };
      ctx.expose({
        open: () => {
          contextInfo.value.img_urls = [];
          contextInfo.value.context = "";
          contextInfo.value.user_addr = "";
          show.value = true;
        },
      });

      return () => (
        <NModal show={show.value} transform-origin="center" auto-focus={false}>
          <OverCard>
            <WrapperHeader>
              <a
                href="#"
                onClick={() => {
                  show.value = false;
                }}
              >
                Cancel
              </a>
              <a href="#" onClick={Done}>
                Done
              </a>
            </WrapperHeader>
            <WrapperBody>
              <NInput
                v-model:value={contextInfo.value.context}
                placeholder="Thoughts at this moment..."
                type="textarea"
                autosize={{
                  minRows: 7,
                  maxRows: 10,
                }}
                size="small"
              ></NInput>
            </WrapperBody>

            <NUpload
              action="/api/upload/file"
              list-type="image-card"
              multiple
              max={9}
              custom-request={customRequest}
              //   onPreview={handlePreview}
            />
            <NModal
              v-model:show={showModalRef.value}
              preset="card"
              style="width: 600px"
            >
              <img src={previewImageUrlRef.value} style="width: 100%" />
            </NModal>
          </OverCard>
        </NModal>
      );
    },
  });
};

export default Write() as any;
