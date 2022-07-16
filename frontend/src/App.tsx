import { useEffect, useState } from "react";
import { SetSelectedCameraMutationVariables, useGetAllCamerasQuery, useSetSelectedCameraMutation } from "./graphql.g";

export const App = () => {
  const { data } = useGetAllCamerasQuery();
  const [selectedCamera, setSelectedCamera] = useState<SetSelectedCameraMutationVariables>();
  const [setSelectedCameraMutation] = useSetSelectedCameraMutation();
  const [isSuccess, setIsSuccess] = useState<boolean | undefined>(undefined);
  useEffect(() => {
    if (!selectedCamera) return;
    (async () => {
      setIsSuccess(undefined);
      const result = await setSelectedCameraMutation({ variables: selectedCamera });
      setIsSuccess(result.data?.setSelectedCamera);
    })();
  }, [selectedCamera]);
  return (
    <>
      <h1>Cameras {selectedCamera && `(${selectedCamera.id} ${selectedCamera.format} ${selectedCamera.frameSize}) => ${isSuccess}`}</h1>
      {data?.cameras.map(({ id, supportedFormats }) =>
        supportedFormats.map(({ format, frameSizes }) =>
          frameSizes.map((frameSize, i) => (
            <div key={id + format + frameSize + i}>
              <button onClick={() => setSelectedCamera({ id, format, frameSize })}>
                {id} _ {format} _ {frameSize}
              </button>
            </div>
          ))
        )
      )}
    </>
  );
};
