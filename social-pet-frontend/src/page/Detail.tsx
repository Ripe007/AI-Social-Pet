import Talk from "@/components/talk";
import axios from "axios";
import { defineComponent, ref, watch } from "vue";
import { useRoute } from "vue-router";
import styled from "vue3-styled-component";

const Detail = () => {
  const Wrapper = styled.div`
    display: flex;
    flex-direction: column;
    width: 100%;
  `;
  return defineComponent({
    setup(props, ctx) {
      const route = useRoute();
      const data = ref<any>(null);
      const init = () => {
        const id = route.query.id;
        axios
          .post("/api/community/list", {
            source_id: id,
          })
          .then((res) => {
            if (res.data.code === 0) {
              data.value = res.data.data[0];
              console.log("data.vlaue", data.value);
            }
          });
      };
      init();
      return () => (
        <Wrapper>
          {data.value && (
            <Talk
              style={{ width: "100%" }}
              userAddr={data.value?.context.user_addr}
              context={data.value?.context.context}
              images={data.value?.context.img_urls}
              source_id={data.value?.context.source_id}
              updated_at={data.value?.context.updated_at}
              like_count={data.value?.config.like_count}
              comment_count={data.value?.config.comment_count}
              forward_count={data.value?.config.forward_count}
            ></Talk>
          )}
        </Wrapper>
      );
    },
  });
};

export default Detail();
