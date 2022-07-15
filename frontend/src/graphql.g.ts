import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Camera = {
  id: Scalars['Int'];
  supportedFormats: Array<SupportedFormat>;
};

export type Mutation = {
  setCamera: Scalars['Boolean'];
};


export type MutationSetCameraArgs = {
  formatName: Scalars['String'];
  frameSize: Scalars['String'];
  id: Scalars['Int'];
};

export type Query = {
  cameras: Array<Camera>;
};

export type SupportedFormat = {
  frameSizes: Array<Scalars['String']>;
  name: Scalars['String'];
};

export type GetAllCamerasQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAllCamerasQuery = { cameras: Array<{ id: number, supportedFormats: Array<{ name: string, frameSizes: Array<string> }> }> };

export type SetCameraMutationVariables = Exact<{
  id: Scalars['Int'];
  formatName: Scalars['String'];
  frameSize: Scalars['String'];
}>;


export type SetCameraMutation = { setCamera: boolean };


export const GetAllCamerasDocument = gql`
    query GetAllCameras {
  cameras {
    id
    supportedFormats {
      name
      frameSizes
    }
  }
}
    `;

/**
 * __useGetAllCamerasQuery__
 *
 * To run a query within a React component, call `useGetAllCamerasQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetAllCamerasQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetAllCamerasQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetAllCamerasQuery(baseOptions?: Apollo.QueryHookOptions<GetAllCamerasQuery, GetAllCamerasQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAllCamerasQuery, GetAllCamerasQueryVariables>(GetAllCamerasDocument, options);
      }
export function useGetAllCamerasLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAllCamerasQuery, GetAllCamerasQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAllCamerasQuery, GetAllCamerasQueryVariables>(GetAllCamerasDocument, options);
        }
export type GetAllCamerasQueryHookResult = ReturnType<typeof useGetAllCamerasQuery>;
export type GetAllCamerasLazyQueryHookResult = ReturnType<typeof useGetAllCamerasLazyQuery>;
export type GetAllCamerasQueryResult = Apollo.QueryResult<GetAllCamerasQuery, GetAllCamerasQueryVariables>;
export const SetCameraDocument = gql`
    mutation SetCamera($id: Int!, $formatName: String!, $frameSize: String!) {
  setCamera(id: $id, formatName: $formatName, frameSize: $frameSize)
}
    `;
export type SetCameraMutationFn = Apollo.MutationFunction<SetCameraMutation, SetCameraMutationVariables>;

/**
 * __useSetCameraMutation__
 *
 * To run a mutation, you first call `useSetCameraMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSetCameraMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [setCameraMutation, { data, loading, error }] = useSetCameraMutation({
 *   variables: {
 *      id: // value for 'id'
 *      formatName: // value for 'formatName'
 *      frameSize: // value for 'frameSize'
 *   },
 * });
 */
export function useSetCameraMutation(baseOptions?: Apollo.MutationHookOptions<SetCameraMutation, SetCameraMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<SetCameraMutation, SetCameraMutationVariables>(SetCameraDocument, options);
      }
export type SetCameraMutationHookResult = ReturnType<typeof useSetCameraMutation>;
export type SetCameraMutationResult = Apollo.MutationResult<SetCameraMutation>;
export type SetCameraMutationOptions = Apollo.BaseMutationOptions<SetCameraMutation, SetCameraMutationVariables>;