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

export type Microphone = {
  name: Scalars['String'];
};

export type Mutation = {
  setSelectedCamera: Scalars['Boolean'];
  setSelectedMicrophone: Scalars['Boolean'];
};


export type MutationSetSelectedCameraArgs = {
  format: Scalars['String'];
  frameSize: Scalars['String'];
  id: Scalars['Int'];
};


export type MutationSetSelectedMicrophoneArgs = {
  name: Scalars['String'];
};

export type Query = {
  audioPath: Scalars['String'];
  cameras: Array<Camera>;
  microphones: Array<Microphone>;
  videoPath: Scalars['String'];
};

export type SupportedFormat = {
  format: Scalars['String'];
  frameSizes: Array<Scalars['String']>;
};

export type GetAllCamerasQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAllCamerasQuery = { cameras: Array<{ id: number, supportedFormats: Array<{ format: string, frameSizes: Array<string> }> }> };

export type GetVideoPathQueryVariables = Exact<{ [key: string]: never; }>;


export type GetVideoPathQuery = { videoPath: string };

export type SetSelectedCameraMutationVariables = Exact<{
  id: Scalars['Int'];
  format: Scalars['String'];
  frameSize: Scalars['String'];
}>;


export type SetSelectedCameraMutation = { setSelectedCamera: boolean };

export type GetAllMicrophonesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAllMicrophonesQuery = { microphones: Array<{ name: string }> };

export type GetAudioPathQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAudioPathQuery = { audioPath: string };

export type SetSelectedMicrophoneMutationVariables = Exact<{
  name: Scalars['String'];
}>;


export type SetSelectedMicrophoneMutation = { setSelectedMicrophone: boolean };


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
export const GetVideoPathDocument = gql`
    query GetVideoPath {
  videoPath
}
    `;
export function useGetVideoPathQuery(baseOptions?: Apollo.QueryHookOptions<GetVideoPathQuery, GetVideoPathQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetVideoPathQuery, GetVideoPathQueryVariables>(GetVideoPathDocument, options);
      }
export function useGetVideoPathLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetVideoPathQuery, GetVideoPathQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetVideoPathQuery, GetVideoPathQueryVariables>(GetVideoPathDocument, options);
        }
export type GetVideoPathQueryHookResult = ReturnType<typeof useGetVideoPathQuery>;
export type GetVideoPathLazyQueryHookResult = ReturnType<typeof useGetVideoPathLazyQuery>;
export const SetSelectedCameraDocument = gql`
    mutation SetSelectedCamera($id: Int!, $format: String!, $frameSize: String!) {
  setSelectedCamera(id: $id, format: $format, frameSize: $frameSize)
}
    `;
export function useSetSelectedCameraMutation(baseOptions?: Apollo.MutationHookOptions<SetSelectedCameraMutation, SetSelectedCameraMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<SetSelectedCameraMutation, SetSelectedCameraMutationVariables>(SetSelectedCameraDocument, options);
      }
export type SetSelectedCameraMutationHookResult = ReturnType<typeof useSetSelectedCameraMutation>;
export const GetAllMicrophonesDocument = gql`
    query GetAllMicrophones {
  microphones {
    name
  }
}
    `;
export function useGetAllMicrophonesQuery(baseOptions?: Apollo.QueryHookOptions<GetAllMicrophonesQuery, GetAllMicrophonesQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAllMicrophonesQuery, GetAllMicrophonesQueryVariables>(GetAllMicrophonesDocument, options);
      }
export function useGetAllMicrophonesLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAllMicrophonesQuery, GetAllMicrophonesQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAllMicrophonesQuery, GetAllMicrophonesQueryVariables>(GetAllMicrophonesDocument, options);
        }
export type GetAllMicrophonesQueryHookResult = ReturnType<typeof useGetAllMicrophonesQuery>;
export type GetAllMicrophonesLazyQueryHookResult = ReturnType<typeof useGetAllMicrophonesLazyQuery>;
export const GetAudioPathDocument = gql`
    query GetAudioPath {
  audioPath
}
    `;
export function useGetAudioPathQuery(baseOptions?: Apollo.QueryHookOptions<GetAudioPathQuery, GetAudioPathQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAudioPathQuery, GetAudioPathQueryVariables>(GetAudioPathDocument, options);
      }
export function useGetAudioPathLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAudioPathQuery, GetAudioPathQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAudioPathQuery, GetAudioPathQueryVariables>(GetAudioPathDocument, options);
        }
export type GetAudioPathQueryHookResult = ReturnType<typeof useGetAudioPathQuery>;
export type GetAudioPathLazyQueryHookResult = ReturnType<typeof useGetAudioPathLazyQuery>;
export const SetSelectedMicrophoneDocument = gql`
    mutation SetSelectedMicrophone($name: String!) {
  setSelectedMicrophone(name: $name)
}
    `;
export function useSetSelectedMicrophoneMutation(baseOptions?: Apollo.MutationHookOptions<SetSelectedMicrophoneMutation, SetSelectedMicrophoneMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<SetSelectedMicrophoneMutation, SetSelectedMicrophoneMutationVariables>(SetSelectedMicrophoneDocument, options);
      }
export type SetSelectedMicrophoneMutationHookResult = ReturnType<typeof useSetSelectedMicrophoneMutation>;