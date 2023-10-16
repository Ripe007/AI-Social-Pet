import { computed, defineComponent, onMounted, ref, watch } from "vue";
import styled from "vue3-styled-component";
import { useWallet } from "../store/wallet";
import { RouterLink, useRoute, useRouter } from "vue-router";
import { mintStore } from "@/store/mint.contract";

const Layout = () => {
  const Wrapper = styled.div`
    display: flex;
    min-height: 100vh;
    flex-direction: column;
  `;
  const WrapperBody = styled.div`
    padding: 12px;
  `;
  const WraperHeader = styled.div`
    display: flex;
    height: 60px;
    width: 100%;
    justify-content: space-between;
    align-items: center;
    padding: 0 12px;
    position: relative;
    h1 {
      font-size: 20px;
      font-weight: bold;
    }
  `;
  const FlexEnd = styled.div`
    display: flex;
    align-items: center;
    justify-content: flex-end;

    img {
      margin-right: 10px;
    }
  `;
  const But = styled.div`
    background: rgb(101 219 181);
    color: #fff;
    padding: 3px 15px;
    border-radius: 5px;
  `;
  const SmallSpan = styled.span`
    font-size: 12px;
  `;

  const MoreItem = styled<any>("div")`
    background: #fff;
    width: 100%;
    position: absolute;
    top: 60px;
    z-index: 1;
    box-shadow: 0px 16px 11px -19px;
    height: ${(props: any) => props.height + "px" || "0px"};
    transition: height 200ms ease-in-out;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  `;
  const ItemAnimate = styled.div`
    border-bottom: 1px solid rgb(226 226 226);
    padding: 5px 16px;
    font-weight: bold;
    text-indent: 2em;
  `;
  return defineComponent({
    setup() {
      const wallet = useWallet();
      const useMintStore = mintStore();
      const isConnect = computed(() => wallet.wallet.isConnect);
      const router = useRouter();
      const route = useRoute();
      const height = ref(0);

      const open = () => {
        height.value = height.value ? 0 : 150;
      };

      const connect = () => {
        wallet.connect();
        localStorage.setItem("connected", "true");
      };
      watch(isConnect, () => {
        useMintStore.init();
      });
      watch(route, () => {
        height.value = 0;
      });

      onMounted(() => {
        const config = localStorage.getItem("connected");
        if (config) {
          wallet.connect();
        }
      });

      return () => (
        <Wrapper>
          <WraperHeader>
            <RouterLink to="/">
              <img src="/assets/logo.png" width={50} alt="" />
            </RouterLink>
            {/* <h1>PurrfectBonds</h1> */}
            <FlexEnd>
              <img
                src="/assets/icon/search.svg"
                width={28}
                alt=""
                onClick={() => {
                  router.push("/");
                }}
              />
              <img src="/assets/icon/book.svg" width={28} alt="" />
              <img
                src="/assets/icon/menu.png"
                width={28}
                alt=""
                onClick={open}
              />
              <But onClick={connect}>
                {wallet.wallet.isConnect ? (
                  <SmallSpan>
                    {wallet.wallet.account.substring(0, 4) +
                      "..." +
                      wallet.wallet.account.substring(38)}
                  </SmallSpan>
                ) : (
                  "Connect"
                )}
              </But>
            </FlexEnd>
          </WraperHeader>
          <MoreItem height={height.value}>
            {height.value ? (
              <>
                <ItemAnimate class="animate__animated animate__fadeInLeft">
                  <RouterLink to="/collection">My Collections</RouterLink>
                </ItemAnimate>
                <ItemAnimate class="animate__animated animate__fadeInRight">
                  <RouterLink to="/community">Community</RouterLink>
                </ItemAnimate>
                <ItemAnimate class="animate__animated animate__fadeInLeft">
                  Marketplace
                </ItemAnimate>
                <ItemAnimate class="animate__animated animate__fadeInRight">
                  Rank
                </ItemAnimate>
              </>
            ) : (
              <div></div>
            )}
          </MoreItem>
          <WrapperBody>
            <router-view />
          </WrapperBody>
        </Wrapper>
      );
    },
  });
};
export default Layout();
