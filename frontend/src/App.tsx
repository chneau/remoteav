import { useEffect, useState } from "react";
import styled from "styled-components";
import { SetSelectedCameraMutationVariables, useGetAllCamerasQuery, useGetVideoPathQuery, useSetSelectedCameraMutation } from "./graphql.g";

const Button = styled.button`
  opacity: 0.2;
  &:hover {
    opacity: 0.8;
  }
`;

interface StreamDivProps {
  stream: string;
}

const StreamDiv = styled.div<StreamDivProps>`
  min-height: 100vh;
  background-image: url("/video");
  background-position: center;
  background-repeat: no-repeat;
  background-size: contain;
  background-color: black;
`;

export const App = () => {
  const { data } = useGetAllCamerasQuery();
  const { data: videoPath } = useGetVideoPathQuery();
  const videoStreamPath = videoPath?.videoPath ?? "";
  const [selectedCamera, setSelectedCamera] = useState<SetSelectedCameraMutationVariables>();
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
    <StreamDiv stream={videoStreamPath}>
      {data?.cameras.map(({ id, supportedFormats }) =>
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
    </StreamDiv>
  );
};
