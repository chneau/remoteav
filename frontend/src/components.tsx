import styled from "styled-components";
import { useGetVideoPathQuery } from "./graphql.g";

export const WhiteText = styled.div`
  color: white;
`;

export const Button = styled.button`
  opacity: 0.15;
  &:hover {
    opacity: 0.6;
  }
`;

const StreamDiv_ = styled.div<{ stream: string }>`
  min-height: 100vh;
  background-image: url(${(x) => x.stream});
  background-position: center;
  background-repeat: no-repeat;
  background-size: contain;
  background-color: black;
`;

export const StreamDiv = () => {
  const { data } = useGetVideoPathQuery();
  if (!data?.videoPath) return null;
  return <StreamDiv_ stream={data.videoPath} />;
};

export const ButtonsDiv = styled.div`
  position: absolute;
  left: 0;
  top: 0;
  display: flex;
  justify-content: space-between;
  width: 100%;
  height: 100%;
  > :last-child {
    text-align: right;
  }
`;
