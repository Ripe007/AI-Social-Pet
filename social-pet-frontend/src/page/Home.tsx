import CollectCard from "@/components/card/home/CollectCard";
import TalkCard from "@/components/card/home/TalkCard";
import { defineComponent } from "vue";
import { RouterLink } from "vue-router";
import styled from "vue3-styled-component";

const Home = () => {
  const Wrapper = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  `;
  const Flex = styled(Wrapper)`
    position: relative;
  `;
  const Bold = styled.span`
    font-size: 30px;
    color: #000;
    font-family: serif;
    font-weight: bold;
  `;
  const PositionBg = styled.div`
    background: rgb(101 219 181);
    width: 80%;
    height: 200px;
    position: absolute;
    z-index: -1;
    bottom: 0;
    transform: rotateZ(45deg);
  `;
  const CollectionText = styled.div`
    font-size: 14px;
    color: #827e7e;
  `;

  const FixedGif = styled.div`
    position: absolute;
    right: 0;
    bottom: 0;
  `;

  const WrapperBodyBg = styled(Wrapper)`
    background: rgb(179 228 212);
    border-radius: 15px;

    p {
      text-align: left;
      width: 100%;
      font-weight: 200;
      font-size: 0.9em;
      text-indent: 2em;
    }
  `;
  const MarketGrop = styled.div`
    background: rgb(144, 210, 193);
    min-height: 140px;
    width: 80%;
    margin: 20px auto;
    border-radius: 15px;
    display: flex;
    align-items: center;
    justify-content: space-around;
  `;
  const RoundFoodImg = styled<any>("img")`
    width: ${(props) => props?.width ?? "100px"};
    height: ${(props) => props?.width ?? "100px"};
    border-radius: 50%;
  `;
  const H2Span = styled(Bold)`
    font-size: 24px;
    align-self: flex-start;
    margin-top: 20px;
    text-indent: 1em;
  `;
  const BoxShadowH2Span = styled(H2Span)`
    display: flex;
    flex-direction: column;
    margin-bottom: 12px;
    &::after {
      content: " ";
      width: 100%;
      height: 5px;
      background: #b3e4d4;
    }
  `;
  const H2SpanNotIndent = styled(H2Span)`
    text-indent: 0;
    font-size: 18px;
    align-self: center;
    border: 1px solid rgb(179 228 212);
    padding: 0 15px;
    border-radius: 10px;
  `;
  const Triangle = styled.div`
    width: 0;
    height: 0;
    border-top: 20px solid rgb(179 228 212);
    border-right: 20px solid transparent;
    border-left: 20px solid transparent;
    margin-bottom: 20px;
  `;

  const Group = styled.div`
    display: flex;
    flex-direction: column;
  `;
  const GroupBotton = styled.div`
    flex: 1;
    background-image: url("/assets/float_bg.svg");
    background-size: 100% 100%;
    text-align:center;
  `;
  const CenterA = styled.a`
    font-size:20px;
    margin:10px auto 20px;
    text-decoration: underline #b3e4d4;
    text-decoration-thickness:3px;
  `
  const StorgeSpan = styled.div`
    width: 70%;
    margin: 20px auto;
    font-weight: bold;
    font-size: 30px;
  `;
  return defineComponent({
    setup() {
      return () => (
        <Wrapper>
          <Flex>
            <Bold>AI.SOCIAL-PET</Bold>
            <CollectionText>Collect and furrever friends</CollectionText>
            <PositionBg />
            <img src="/assets/banner1.png" />
            <FixedGif>
              <img src="/assets/cure1.gif" width={100} alt="" />
            </FixedGif>
          </Flex>
          <WrapperBodyBg>
            <CollectCard></CollectCard>
            <H2Span>MARKET</H2Span>
            <p style={{ width: "80%", textIndent: 0 }}>
              1.Buy your favorite pet at the market
            </p>
            <p style={{ width: "80%", textIndent: 0 }}>
              2.Anyone can trade on it
            </p>
            <p style={{ textIndent: 0, width: "80%" }}>
              3.Half of every transaction tax goes into the foundation, which
              will be used to assist injured wild animals and pet owners in need
              of pet rescue funds.
            </p>
            <MarketGrop>
              <RoundFoodImg src="/assets/market1.png" width={100} alt="" />
              <RoundFoodImg src="/assets/food.png" width={100} />
            </MarketGrop>
          </WrapperBodyBg>
          <Triangle />
          <Flex>
            <RoundFoodImg src="/assets/market.png" width={100} alt="" />
            <H2SpanNotIndent>To Market</H2SpanNotIndent>
          </Flex>
          <BoxShadowH2Span>Foundation</BoxShadowH2Span>
          <Group>
            <img src="/assets/straycat.png" alt="" />
            <GroupBotton>
              <StorgeSpan>
                Saving Wildlife, Supporting Pet Lovers, One Heart at a Time
              </StorgeSpan>
              <RouterLink to="/foundation">
                <CenterA href="#">View</CenterA>
              </RouterLink>
            </GroupBotton>
          </Group>

          <BoxShadowH2Span>Community</BoxShadowH2Span>
          <TalkCard style={{ margin: "20px auto" }} />
        </Wrapper>
      );
    },
  });
};
export default Home();
