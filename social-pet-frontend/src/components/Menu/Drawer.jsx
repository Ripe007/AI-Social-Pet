import {
  NButton,
  NCard,
  NDrawer,
  NInput,
  NModal,
  NSpin,
  useMessage,
} from "naive-ui";
import { computed, defineComponent, onMounted } from "vue";
import { useRouter } from "vue-router";
import styled from "vue3-styled-component";
import { HomeCard } from "../card/layout";
import { useMiner } from "@/store/miner";
import { useI18n } from "vue-i18n";

const Drawer = styled(NDrawer)`
  .n-drawer.n-drawer--right-placement {
    bottom: 50vh;
  }
`;
const Card = styled(NCard)`
  width: 90vw;
  border-radius: 15px;
`;
const Group = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
`;
const Content = styled.div`
  width: ${(props) => props.width ?? 0}px;
  height: ${(props) => props.height ?? 0}px;
  background-image: linear-gradient(135deg, #4660e5, #7e00ff);
  position: absolute;
  top: 0;
  right: 0;
  border-radius: 0 0 0 100%;
  display: flex;
  justify-content: center;

  ul {
    color: #fff;
    display: ${(props) => (props.width ? "block" : "none")};
  }
  li {
    margin: 16px 0;
    list-style: none;
    font-size: 1.2em;
    color: rgba(255, 255, 255, 0.7);
    font-weight: bold;
  }
  a {
    color: #fff;
  }
`;
const BoldText = styled.div`
  font-size: 1.2em;
  font-weight: bold;
  color: #000;
`;
const Bind = defineComponent({
  setup() {
    const { t } = useI18n();
    const inviteVal = ref(null);
    const minerStore = useMiner();
    const message = useMessage();
    const loading = ref(false);
    const isShow = computed(() => {
      if (minerStore.info.parent === null) return false;
      let bool = !Number(minerStore.info.parent);
      return bool;
    });
    const errorTip = ref(false);
    const onBind = () => {
      loading.value = true;
      minerStore
        .bind(inviteVal.value)
        .then((res) => {
          message.success(t("home.bindsuccessful"));
        })
        .catch((err) => {
          message.error(err);
        })
        .finally(() => {
          loading.value = false;
        });
    };
    onMounted(async () => {
      let val = window.location.href.indexOf("?code=");
      if (val !== -1) {
        let parent = window.location.href.substring(val + 6, val + 48);
        if (parent === "0x758996438C82ef46991938139687037f72F3A5d8") {
          inviteVal.value = parent;
        } else {
          const bool = (await minerStore.getInv(parent)).gte(10e18);
          if (bool) {
            inviteVal.value = parent;
          } else {
            errorTip.value = !bool;
          }
        }
      }
    });
    minerStore.initstatic();
    return () => (
      <>
        <NModal v-model:show={isShow.value} transform-origin="center">
          <Card>
            <BoldText>{t("home.bind")}</BoldText>
            <NInput
              style="margin:16px 0;"
              placeholder={t("home.useLink")}
              readonly
              value={inviteVal.value}
            />
            {errorTip.value && (
              <div style="color:#d81e06;text-align:center;">
                {t("home.useOtherLink")}
              </div>
            )}
            <Group>
              <NButton
                color="#6814CC"
                loading={loading.value}
                disabled={!inviteVal.value}
                onClick={onBind}
                style="min-width:100px;"
                round
              >
                {t("confirm")}
              </NButton>
            </Group>
          </Card>
        </NModal>
      </>
    );
  },
});
export default defineComponent({
  setup(props, { expose }) {
    const { t } = useI18n();
    const isShow = ref(false);
    const router = useRouter();
    const toHref = (path) => {
      router.push({ path });
      isShow.value = false;
    };
    expose({
      open: () => (isShow.value = true),
    });
    return () => (
      <>
        <Drawer v-model:show={isShow.value} placement="right" width="0">
          <Content
            width={isShow.value && 300}
            height={isShow.value && 300}
            class="animate__animated animate__slideInRight"
          >
            <ul>
              <li
                class="animate__animated animate__bounceInLeft"
                onClick={toHref.bind(null, "/")}
              >
                {t("router.home")}
              </li>
              <li
                class="animate__animated animate__bounceInRight"
                onClick={toHref.bind(null, "/data")}
              >
                {t("router.data")}
              </li>
              <li
                class="animate__animated animate__bounceInLeft"
                onClick={toHref.bind(null, "/myAssets")}
              >
                {t("router.assets")}
              </li>
              <li
                class="animate__animated animate__bounceInRight"
                onClick={toHref.bind(null, "/share")}
              >
                {t("router.share")}
              </li>
            </ul>
          </Content>
        </Drawer>
        <Bind />
      </>
    );
  },
});
