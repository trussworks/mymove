import React, { useState, useEffect, useMemo } from 'react';
import { GridContainer } from '@trussworks/react-uswds';
import { useTable, useFilters, usePagination } from 'react-table';
import PropTypes from 'prop-types';

import styles from './TableQueue.module.scss';

import Table from 'components/Table/Table';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import TextBoxFilter from 'components/Table/Filters/TextBoxFilter';

// TableQueue is a react-table that uses react-hooks to fetch, filter, sort and page data
const TableQueue = ({ title, columns, handleClick, useQueries }) => {
  const [paramFilters, setParamFilters] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [currentPageSize, setCurrentPageSize] = useState(20);
  const [pageCount, setPageCount] = React.useState(0);

  const {
    queueResult: { totalCount = 0, data = [], page = 1, perPage = 20 },
    isLoading,
    isError,
  } = useQueries({ filters: paramFilters, currentPage, currentPageSize });

  // react-table setup below

  const defaultColumn = useMemo(
    () => ({
      // Let's set up our default Filter UI
      Filter: TextBoxFilter,
    }),
    [],
  );
  const tableData = useMemo(() => data, [data]);
  const tableColumns = useMemo(() => columns, [columns]);
  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow,
    canPreviousPage,
    canNextPage,
    pageOptions,
    gotoPage,
    nextPage,
    previousPage,
    setPageSize,
    state: { filters, pageIndex, pageSize },
  } = useTable(
    {
      columns: tableColumns,
      data: tableData,
      initialState: { hiddenColumns: ['id'], pageSize: perPage, pageIndex: page - 1 },
      defaultColumn, // Be sure to pass the defaultColumn option
      manualFilters: true,
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

  return (
    <GridContainer containerSize="widescreen" className={styles.TableQueue}>
      <h1>{`${title} (${totalCount})`}</h1>
      <div className={styles.tableContainer}>
        <Table
          showPagination
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

TableQueue.propTypes = {
  // handleClick is the handler to handle functionality to click on a row
  handleClick: PropTypes.func.isRequired,
  // useQueries is the react-query hook call to handle data fetching
  useQueries: PropTypes.func.isRequired,
  // title is the table title
  title: PropTypes.string.isRequired,
  // columns is the columns to show in the table
  columns: PropTypes.arrayOf(PropTypes.object).isRequired,
};

export default TableQueue;
