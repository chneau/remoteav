import { useEffect, useState } from "react";
import { SetCameraMutationVariables, useGetAllCamerasQuery, useSetCameraMutation } from "./graphql.g";

export const App = () => {
  const { data } = useGetAllCamerasQuery();
  const [selected, setSelected] = useState<SetCameraMutationVariables>();
  const [setCameraMutation] = useSetCameraMutation();
  useEffect(() => {
    if (!selected) return;
    setCameraMutation({ variables: selected });
  }, [selected]);
  return (
    <>
      <h1>Cameras {selected && `(${selected.id} ${selected.format} ${selected.frameSize})`}</h1>
      {data?.cameras.map(({ id, supportedFormats }) =>
        supportedFormats.map(({ format, frameSizes }) =>
          frameSizes.map((frameSize, i) => (
            <div key={id + format + frameSize + i}>
              <button onClick={() => setSelected({ id, format, frameSize })}>
                {id} _ {format} _ {frameSize}
              </button>
            </div>
          ))
        )
      )}
    </>
  );
};
