import { useGetCameraIdsQuery } from "./generated/graphql";

export const App = () => {
  const { data } = useGetCameraIdsQuery();

  return <pre>{JSON.stringify(data, null, 2)}</pre>;
};
