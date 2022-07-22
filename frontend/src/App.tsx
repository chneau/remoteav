import { useEffect, useState } from "react";
import { Button, ButtonsDiv, StreamDiv, WhiteText } from "./components";
import {
  SetSelectedCameraMutationVariables,
  useGetAllCamerasQuery,
  useGetAllMicrophonesQuery,
  useSetSelectedCameraMutation,
  useSetSelectedMicrophoneMutation,
} from "./graphql.g";

export const App = () => {
  const [selectedCamera, setSelectedCamera] = useState<SetSelectedCameraMutationVariables>();
  const [selectedMicrophone, setSelectedMicrophone] = useState<string>();
  return (
    <>
      <StreamDiv />
      <ButtonsDiv>
        <CameraSelector onSelectedCamera={(x) => setSelectedCamera(x)} />
        <WhiteText>
          Video: {selectedCamera ? `${selectedCamera.id} _ ${selectedCamera.format} _ ${selectedCamera.frameSize}` : "none"}
          <br />
          Audio: {selectedMicrophone || "none"}
        </WhiteText>
        <MicrophoneSelector onSelectedMicrophone={(x) => setSelectedMicrophone(x)} />
      </ButtonsDiv>
    </>
  );
};

const MicrophoneSelector = ({ onSelectedMicrophone }: { onSelectedMicrophone?: (x: string) => any }) => {
  const { data: microphones } = useGetAllMicrophonesQuery();
  const [selectedMicrophone, setSelectedMicrophone] = useState<string>();
  const [setSelectedMicrophoneMutation] = useSetSelectedMicrophoneMutation();
  useEffect(() => {
    if (!selectedMicrophone) return;
    onSelectedMicrophone?.(selectedMicrophone);
    (async () => console.log((await setSelectedMicrophoneMutation({ variables: { name: selectedMicrophone } })).data))();
  }, [selectedMicrophone]);
  return (
    <div>
      {microphones?.microphones.map(({ name }, i) => (
        <div key={name + i}>
          <Button onClick={() => setSelectedMicrophone(name)}>{name}</Button>
        </div>
      ))}
    </div>
  );
};

const CameraSelector = ({ onSelectedCamera }: { onSelectedCamera?: (x: SetSelectedCameraMutationVariables) => any }) => {
  const { data: cameras } = useGetAllCamerasQuery();
  const [selectedCamera, setSelectedCamera] = useState<SetSelectedCameraMutationVariables>();
  const [setSelectedCameraMutation] = useSetSelectedCameraMutation();
  useEffect(() => {
    if (!selectedCamera) return;
    onSelectedCamera?.(selectedCamera);
    (async () => console.log((await setSelectedCameraMutation({ variables: selectedCamera })).data))();
    document.title = selectedCamera ? `${selectedCamera.id} ${selectedCamera.format} ${selectedCamera.frameSize}` : "";
  }, [JSON.stringify(selectedCamera)]);
  return (
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
  );
};
