import React, { useState, useEffect, useMemo } from 'react';
import { withRouter } from 'react-router-dom';
import { GridContainer } from '@trussworks/react-uswds';
import { useTable, useFilters, usePagination } from 'react-table';

import styles from './MoveQueue.module.scss';

import { HistoryShape } from 'types/router';
import Table from 'components/Table/Table';
import { createHeader } from 'components/Table/utils';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import { useMovesQueueQueries } from 'hooks/queries';
import { serviceMemberAgencyLabel } from 'shared/formatters';
import TextBoxFilter from 'components/Table/Filters/TextBoxFilter';
import MultiSelectCheckBoxFilter from 'components/Table/Filters/MultiSelectCheckBoxFilter';
import SelectFilter from 'components/Table/Filters/SelectFilter';
import { BRANCH_OPTIONS, MOVE_STATUS_OPTIONS } from 'constants/queues';

const moveStatusOptions = Object.keys(MOVE_STATUS_OPTIONS).map((key) => ({
  value: key,
  label: MOVE_STATUS_OPTIONS[`${key}`],
}));

const branchFilterOptions = [
  { value: '', label: 'All' },
  ...Object.keys(BRANCH_OPTIONS).map((key) => ({
    value: key,
    label: BRANCH_OPTIONS[`${key}`],
  })),
];

const columns = [
  createHeader('ID', 'id'),
  createHeader(
    'Customer name',
    (row) => {
      return `${row.customer.last_name}, ${row.customer.first_name}`;
    },
    {
      id: 'lastName',
      isFilterable: true,
    },
  ),
  createHeader('DoD ID', 'customer.dodID', {
    id: 'dodID',
    isFilterable: true,
  }),
  createHeader('Status', 'status', {
    isFilterable: true,
    // eslint-disable-next-line react/jsx-props-no-spreading
    Filter: (props) => <MultiSelectCheckBoxFilter options={moveStatusOptions} {...props} />,
  }),
  createHeader('Move Code', 'locator', {
    id: 'moveID',
    isFilterable: true,
  }),
  createHeader(
    'Branch',
    (row) => {
      return serviceMemberAgencyLabel(row.customer.agency);
    },
    {
      id: 'branch',
      isFilterable: true,
      // eslint-disable-next-line react/jsx-props-no-spreading
      Filter: (props) => <SelectFilter options={branchFilterOptions} {...props} />,
    },
  ),
  createHeader('# of shipments', 'shipmentsCount'),
  createHeader('Destination duty station', 'destinationDutyStation.name', {
    id: 'destinationDutyStation',
    isFilterable: true,
  }),
  createHeader('Origin GBLOC', 'originGBLOC'),
];

const MoveQueue = ({ history }) => {
  const [paramFilters, setParamFilters] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [currentPageSize, setCurrentPageSize] = useState(20);
  const [pageCount, setPageCount] = React.useState(0);

  const {
    queueMovesResult: { totalCount = 0, queueMoves = [], page = 1, perPage = 20 },
    isLoading,
    isError,
  } = useMovesQueueQueries({ filters: paramFilters, currentPage, currentPageSize });

  // react-table setup below

  const defaultColumn = useMemo(
    () => ({
      // Let's set up our default Filter UI
      Filter: TextBoxFilter,
    }),
    [],
  );
  const tableData = useMemo(() => queueMoves, [queueMoves]);
  const tableColumns = useMemo(() => columns, []);
  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow,
    canPreviousPage,
    canNextPage,
    gotoPage,
    pageOptions,
    previousPage,
    nextPage,
    setPageSize,
    state: { filters, pageIndex, pageSize },
  } = useTable(
    {
      columns: tableColumns,
      data: tableData,
      initialState: { hiddenColumns: ['id'], pageSize: perPage, pageIndex: page - 1 },
      defaultColumn, // Be sure to pass the defaultColumn option
      manualFilters: true,
      showPagination: true,
      manualPagination: true,
      pageCount,
    },
    useFilters,
    usePagination,
  );

  // When these table states change, fetch new data!
  useEffect(() => {
    if (!isLoading && !isError) {
      setParamFilters(filters);
      setCurrentPage(pageIndex + 1);
      setCurrentPageSize(pageSize);
      setPageCount(Math.ceil(totalCount / pageSize));
    }
  }, [filters, pageIndex, pageSize, isLoading, isError, totalCount]);

  if (isLoading) return <LoadingPlaceholder />;
  if (isError) return <SomethingWentWrong />;

  const handleClick = (values) => {
    history.push(`/moves/${values.id}/details`);
  };

  return (
    <GridContainer containerSize="widescreen" className={styles.MoveQueue}>
      <h1>{`All moves (${totalCount})`}</h1>
      <div className={styles.tableContainer}>
        <Table
          handleClick={handleClick}
          gotoPage={gotoPage}
          setPageSize={setPageSize}
          nextPage={nextPage}
          previousPage={previousPage}
          getTableProps={getTableProps}
          getTableBodyProps={getTableBodyProps}
          headerGroups={headerGroups}
          rows={rows}
          prepareRow={prepareRow}
          showPagination
          canPreviousPage={canPreviousPage}
          canNextPage={canNextPage}
          pageIndex={pageIndex}
          pageSize={pageSize}
          pageCount={pageCount}
          pageOptions={pageOptions}
        />
      </div>
    </GridContainer>
  );
};

MoveQueue.propTypes = {
  history: HistoryShape.isRequired,
};

export default withRouter(MoveQueue);
