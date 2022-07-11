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

export type Query = {
  cameras: Array<Camera>;
};

export type SupportedFormat = {
  format: Scalars['String'];
  frameSizes: Array<Scalars['String']>;
};

export type GetCameraIdsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetCameraIdsQuery = { cameras: Array<never> };


export const GetCameraIdsDocument = gql`
    query GetCameraIds {
  cameras {
    id
    supportedFormats {
      format
      frameSizes
    }
  }
}
    `;

/**
 * __useGetCameraIdsQuery__
 *
 * To run a query within a React component, call `useGetCameraIdsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetCameraIdsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetCameraIdsQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetCameraIdsQuery(baseOptions?: Apollo.QueryHookOptions<GetCameraIdsQuery, GetCameraIdsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetCameraIdsQuery, GetCameraIdsQueryVariables>(GetCameraIdsDocument, options);
      }
export function useGetCameraIdsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetCameraIdsQuery, GetCameraIdsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetCameraIdsQuery, GetCameraIdsQueryVariables>(GetCameraIdsDocument, options);
        }
export type GetCameraIdsQueryHookResult = ReturnType<typeof useGetCameraIdsQuery>;
export type GetCameraIdsLazyQueryHookResult = ReturnType<typeof useGetCameraIdsLazyQuery>;
export type GetCameraIdsQueryResult = Apollo.QueryResult<GetCameraIdsQuery, GetCameraIdsQueryVariables>;