import { Block } from 'baseui/block'
import { gql, useQuery } from '@apollo/client'
import { Container } from '../../Components/Layout'
import { NavBar } from '../../Components/Navbar'
import { Item, ITEM_FIELDS } from '../../Components/Item'
import { useSearchParams } from 'react-router-dom'
import { PaddedBlock } from '../../Components/Layout'
import { LoadingScreen } from './LoadingScreen'
import { ErrorScreen } from './ErrorScreen'
import { Job, NodeT, Story } from '../../types'
import { Fragment } from 'react'

const GET_ITEM_QUERY = gql`
  ${ITEM_FIELDS}
  query GetItem($id: ID!) {
    item: node(id: $id) {
      ...ItemFields
      ...JobFields
    }
  }
`

interface GetItemQueryData {
  item: NodeT<Story | Job>
}

interface GetItemQueryVars {
  id: string
}

const ItemPage: React.FC = () => {
  const [searchParams] = useSearchParams()
  const { data, loading, error } = useQuery<GetItemQueryData, GetItemQueryVars>(
    GET_ITEM_QUERY,
    {
      variables: {
        id: searchParams.get('id') || '',
      },
    }
  )
  let children: React.ReactNode = <Fragment />
  if (!data && loading) {
    children = <LoadingScreen />
  } else if (!data && error) {
    children = <ErrorScreen error={error} />
  } else if (data) {
    children = <Item item={data.item} />
  }
  return (
    <Container
      top={<NavBar />}
      left={<Block />}
      center={<PaddedBlock>{children}</PaddedBlock>}
      right={<Block />}
    />
  )
}

export { ItemPage }
