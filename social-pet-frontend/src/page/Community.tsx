import Write from "@/components/card/community/Write";
import Talk from "@/components/talk";
import axios from "axios";
import { defineComponent, inject, ref } from "vue";
import styled from "vue3-styled-component";

const Community = () => {
  const Wrapper = styled.div`
    margin-top: 10%;
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    position: relative;
  `;
  const TitleGroup = styled.div`
    width: 100%;
    line-height: 1;
    border-bottom: 1px solid #e6dbdb;
    display: flex;
    align-items: center;
  `;
  const Title = styled<any>("div")`
    width: 50%;
    height: 2em;
    line-height: 2em;
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;
    &:active {
      background: #f9f9f9;
    }

    &::after {
      display: ${(props) => (props.active ? "block" : "none")};
      content: " ";
      width: 80px;
      height: 3px;
      background: rgb(101 219 181);
      margin-top: -5px;
    }
  `;

  const MineBtnGroup = styled.div`
    position: absolute;
    display: flex;
    justify-content: flex-end;
    top: -45px;
    right: 20px;

    img {
      margin: 0 5px;
      &:active {
        background: #f9f9f9;
      }
    }
  `;
  return defineComponent({
    setup(props, ctx) {
      const activeType = ref(0);
      const writeRef = ref<any>(null);
      const contexts = ref<any>([]);
      const pageLoader = inject("pageLoader", (bool: boolean) => {});
      const change = (type: number) => {
        activeType.value = type;
      };

      const init = () => {
        pageLoader(true);
        axios.post("/api/community/list").then((res) => {
          contexts.value = res.data.data;
          setTimeout(() => {
            pageLoader(false);
          }, 500);
        });
      };
      init();
      const openWrite = () => {
        writeRef.value?.open();
      };
      return () => (
        <Wrapper>
          <MineBtnGroup>
            <img
              src="/assets/icon/write.svg"
              width={30}
              alt=""
              onClick={openWrite}
            />
            <img src="/assets/icon/mine.svg" width={30} alt="" />
          </MineBtnGroup>
          <TitleGroup>
            <Title
              active={activeType.value === 0}
              onClick={change.bind(null, 0)}
            >
              Hot
            </Title>
            <Title
              active={activeType.value === 1}
              onClick={change.bind(null, 1)}
            >
              Latest
            </Title>
          </TitleGroup>
          <Write ref={writeRef} />
          <div style={{ width: "100%" }}>
            {contexts.value.map((item: any) => (
              <Talk
                key={item.context.source_id}
                userAddr={item.context.user_addr}
                context={item.context.context}
                images={item.context.img_urls}
                source_id={item.context.source_id}
                updated_at={item.context.updated_at}
                like_count={item.config.like_count}
                comment_count={item.config.comment_count}
                forward_count={item.config.forward_count}
              ></Talk>
            ))}
          </div>
        </Wrapper>
      );
    },
  });
};

export default Community();
