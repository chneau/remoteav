import { useEffect, useState } from "react";
import styled from "styled-components";
import {
  SetSelectedCameraMutationVariables,
  useGetAllCamerasQuery,
  useGetAllMicrophonesQuery,
  useGetVideoPathQuery,
  useSetSelectedCameraMutation,
} from "./graphql.g";

const Button = styled.button`
  opacity: 0.15;
  &:hover {
    opacity: 0.6;
  }
`;

const StreamDiv = styled.div<{ stream: string }>`
  min-height: 100vh;
  background-image: url(${(x) => x.stream});
  background-position: center;
  background-repeat: no-repeat;
  background-size: contain;
  background-color: black;
`;

const ButtonsDiv = styled.div`
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

export const App = () => {
  const { data: cameras } = useGetAllCamerasQuery();
  const { data: videoPath } = useGetVideoPathQuery();
  const { data: microphones } = useGetAllMicrophonesQuery();
  const videoStreamPath = videoPath?.videoPath ?? "";
  const [selectedCamera, setSelectedCamera] = useState<SetSelectedCameraMutationVariables>();
  const [selectedMicrophone, setSelectedMicrophone] = useState<string>();
  const [setSelectedCameraMutation] = useSetSelectedCameraMutation();
  const [isSuccess, setIsSuccess] = useState<boolean | undefined>(undefined);
  const selectedCameraText = (selectedCamera && `(${selectedCamera.id} ${selectedCamera.format} ${selectedCamera.frameSize}) => ${isSuccess}`) || "";
  useEffect(() => {
    if (!selectedCamera) return;
    (async () => {
      setIsSuccess(undefined);
      const result = await setSelectedCameraMutation({ variables: selectedCamera });
      setIsSuccess(result.data?.setSelectedCamera);
    })();
    document.title = selectedCameraText;
  }, [selectedCamera]);
  return (
    <>
      <StreamDiv stream={videoStreamPath} />
      <ButtonsDiv>
        <div>
          {cameras?.cameras.map(({ id, supportedFormats }) =>
            supportedFormats.map(({ format, frameSizes }) =>
              frameSizes.map((frameSize, i) => (
                <div key={id + format + frameSize + i}>
                  <Button onClick={() => setSelectedCamera({ id, format, frameSize })}>
                    {id} _ {format} _ {frameSize}
                  </Button>
                </div>
              ))
            )
          )}
        </div>
        <div>
          {microphones?.microphones.map(({ name }, i) => (
            <div>
              <Button key={name + i} onClick={() => setSelectedMicrophone(name)}>
                {name}
              </Button>
            </div>
          ))}
        </div>
      </ButtonsDiv>
    </>
  );
};
