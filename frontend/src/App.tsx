import { useState } from "react";
import { useGetAllCamerasQuery } from "./graphql.g";

interface SelectedCamera {
  id: number;
  name: string;
  frameSize: string;
}

export const App = () => {
  const { data } = useGetAllCamerasQuery();
  const [selected, setSelected] = useState<SelectedCamera>();
  return (
    <>
      <h1>Cameras {selected && `(${selected.id} ${selected.name} ${selected.frameSize})`}</h1>
      {data?.cameras.map(({ id, supportedFormats }) =>
        supportedFormats.map(({ name, frameSizes }) =>
          frameSizes.map((frameSize, i) => (
            <div key={id + name + frameSize + i}>
              <button onClick={() => setSelected({ id, name, frameSize })}>
                {id} _ {name} _ {frameSize}
              </button>
            </div>
          ))
        )
      )}
    </>
  );
};
