import { useGetAllCamerasQuery } from "./graphql.g";

export const App = () => {
  const { data } = useGetAllCamerasQuery();
  const cameras = data?.cameras?.filter((x) => x.supportedFormats?.length > 0);
  return <pre>{JSON.stringify(cameras, null, 2)}</pre>;
};
