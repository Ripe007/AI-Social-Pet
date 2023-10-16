import { useWallet } from "@/store/wallet";
import BigNumber from "bignumber.js";
import { computed, defineComponent, onMounted, ref, watch } from "vue";
import styled from "vue3-styled-component";

const Foundation = () => {
  const Wrapper = styled.div`
    margin-top: 10%;
    width: 100%;
    min-height: 80vh;
    box-shadow: 0px 0px 10px 0px #e6dbdb;
    border-radius: 15px;
    background: #fff;
    padding: 30px 16px 16px;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    position: relative;
  `;
  const RoundIcon = styled.img`
    width: 80px;
    height: 80px;
    border-radius: 50%;
    box-shadow: 0px 0px 10px 0px #e6dbdb;
    position: absolute;
    top: -30px;
  `;
  const Bold = styled.span`
    font-size: 30px;
    color: #000;
    font-family: serif;
    font-weight: bold;
  `;
  const H2Span = styled(Bold)`
    font-size: 24px;
    margin-top: 20px;
    margin-right: 10px;
  `;
  const WrapperBody = styled.div`
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    width: 100%;
    margin-top: 16px;

    ul {
      margin-bottom: 20px;
    }
    li {
      font-size: 16px;
      margin-top: 10px;
    }
  `;
  const Title = styled.div`
    font-size: 18px;
    font-weight: bold;
  `;
  return defineComponent({
    setup() {
      const wallet = useWallet();
      const foundationBal = ref(BigNumber(0));

      onMounted(() => {
        wallet.getBalanceOf((import.meta as any).VITE_FONDATION).then((val) => {
          if (val) {
            foundationBal.value = val;
          }
        });
      });
      return () => (
        <Wrapper>
          <RoundIcon src="/assets/loader.gif" />

          <div style={{ marginTop: "30px" }}>
            <H2Span>{foundationBal.value.div(1e18).toFixed(4)}</H2Span>
            eth
          </div>
          <div style={{ marginTop: "10px", fontSize: "18px" }}>
            Foundation Pool
          </div>

          <WrapperBody>
            <Title>Fund source</Title>
            <ul>
              <li>1.User coinage income</li>
              <li>2.The tax on each transaction in the trading market</li>
              <li>3.Donations from loving people</li>
            </ul>

            <Title>Fund expense</Title>
            <ul>
              <li>1.Rescue wild animals</li>
              <li>
                2.Help people with financial difficulties who need pet therapy
              </li>
              <li>3.Donations from loving people</li>
            </ul>

            <p>
              * Coinage users have voting rights, and each time they spend, they
              need to publicly receive more than half of the voting rights
            </p>
          </WrapperBody>
        </Wrapper>
      );
    },
  });
};

export default Foundation;
