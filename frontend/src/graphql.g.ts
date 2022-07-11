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
  __typename?: 'Query';
  cameras: Array<Camera>;
};

export type SupportedFormat = {
  format: Scalars['String'];
  frameSizes: Array<Scalars['String']>;
};

export type GetAllCamerasQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAllCamerasQuery = { __typename?: 'Query', cameras: Array<never> };


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