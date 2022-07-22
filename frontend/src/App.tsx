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
  const { data: cameras } = useGetAllCamerasQuery();
  const { data: microphones } = useGetAllMicrophonesQuery();
  const [selectedCamera, setSelectedCamera] = useState<SetSelectedCameraMutationVariables>();
  const [selectedMicrophone, setSelectedMicrophone] = useState<string>();
  const [setSelectedCameraMutation] = useSetSelectedCameraMutation();
  const [setSelectedMicrophoneMutation] = useSetSelectedMicrophoneMutation();
  useEffect(() => {
    if (!selectedMicrophone) return;
    (async () => console.log((await setSelectedMicrophoneMutation({ variables: { name: selectedMicrophone } })).data))();
  }, [selectedMicrophone]);

  useEffect(() => {
    if (!selectedCamera) return;
    (async () => console.log((await setSelectedCameraMutation({ variables: selectedCamera })).data))();
    document.title = selectedCamera ? `${selectedCamera.id} ${selectedCamera.format} ${selectedCamera.frameSize}` : "";
  }, [JSON.stringify(selectedCamera)]);
  return (
    <>
      <StreamDiv />
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
        <WhiteText>
          Video: {selectedCamera ? `${selectedCamera.id} _ ${selectedCamera.format} _ ${selectedCamera.frameSize}` : "none"}
          <br />
          Audio: {selectedMicrophone || "none"}
        </WhiteText>
        <div>
          {microphones?.microphones.map(({ name }, i) => (
            <div key={name + i}>
              <Button onClick={() => setSelectedMicrophone(name)}>{name}</Button>
            </div>
          ))}
        </div>
      </ButtonsDiv>
    </>
  );
};
