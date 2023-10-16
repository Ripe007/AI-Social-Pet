import { Transition, defineComponent, onMounted, provide, ref } from "vue";
import Layout from "./layout";
import styled from "vue3-styled-component";
import { NMessageProvider, NModal } from "naive-ui";

const App = () => {
  const LoaderWrapper = styled.div`
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
    // background:trans
    img {
      width: 100px;
    }
  `;

  return defineComponent({
    setup() {
      const pageLoading = ref(false);
      const loader = ref(true);
      provide("pageLoader", (status: boolean) => {
        pageLoading.value = status;
      });

      onMounted(() => {
        setTimeout(() => {
          loader.value = false;
        }, Math.random() * 2000);
      });
      return () => (
        <>
          <Transition>
            {loader.value && (
              <LoaderWrapper>
                <img src="/assets/loader.gif" alt="" />
              </LoaderWrapper>
            )}
            <NMessageProvider>
              <Layout />
            </NMessageProvider>
          </Transition>
          <NModal
            z-index={9999}
            show={pageLoading.value}
            transformOrigin="center"
          >
            <img src="/assets/loader.gif" width={100} alt="" />
          </NModal>
        </>
      );
    },
  });
};
export default App();
