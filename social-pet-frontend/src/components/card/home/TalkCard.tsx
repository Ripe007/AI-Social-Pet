import Talk from "@/components/talk";
import axios from "axios";
import { defineComponent, ref } from "vue";
import { RouterLink } from "vue-router";
import styled from "vue3-styled-component";

const TalkCard = () => {
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
  return defineComponent({
    setup() {
      const contexts = ref([]);
      const init = () => {
        axios.post("/api/community/list").then((res) => {
          if (res.data.data.length > 3) {
            contexts.value = res.data.data.slice(0, 3);
          } else {
            contexts.value = res.data.data;
          }
        });
      };
      init();
      return () => (
        <Card>
          {contexts.value.map((item: any) => (
            <Talk
              style={{ width: "100%" }}
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
          <RouterLink to="/community">More...</RouterLink>
        </Card>
      );
    },
  });
};

export default TalkCard();
