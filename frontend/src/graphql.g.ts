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
  format: Scalars['String'];
  frameSize: Scalars['String'];
  id: Scalars['Int'];
};

export type Query = {
  cameras: Array<Camera>;
};

export type SupportedFormat = {
  format: Scalars['String'];
  frameSizes: Array<Scalars['String']>;
};

export type GetAllCamerasQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAllCamerasQuery = { cameras: Array<{ id: number, supportedFormats: Array<{ format: string, frameSizes: Array<string> }> }> };

export type SetCameraMutationVariables = Exact<{
  id: Scalars['Int'];
  format: Scalars['String'];
  frameSize: Scalars['String'];
}>;


export type SetCameraMutation = { setCamera: boolean };


export const GetAllCamerasDocument = gql`
    query GetAllCameras {
  cameras {
    id
    supportedFormats {
      format
      frameSizes
    }
  }
}
    `;
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
export const SetCameraDocument = gql`
    mutation SetCamera($id: Int!, $format: String!, $frameSize: String!) {
  setCamera(id: $id, format: $format, frameSize: $frameSize)
}
    `;
export function useSetCameraMutation(baseOptions?: Apollo.MutationHookOptions<SetCameraMutation, SetCameraMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<SetCameraMutation, SetCameraMutationVariables>(SetCameraDocument, options);
      }
export type SetCameraMutationHookResult = ReturnType<typeof useSetCameraMutation>;