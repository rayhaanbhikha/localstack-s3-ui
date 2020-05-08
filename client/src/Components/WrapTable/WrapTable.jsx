import React from 'react'
import './styles.css'

import { RefreshIcon } from '..'


export const WrapTable = (TableText, TableBody, refreshResources) => {
  return (
    <div className="buckets-table">
      <div className="table-head-container">
        <div className="table-text">
          <TableText />
        </div>
        <div style={{ "alignSelf": "flex-end" }} onClick={refreshResources}>
          <RefreshIcon />
        </div>
      </div>
      <table>
        <thead >
          <tr className="table-column-heading">
            <th className="table-column-heading-text">Name</th>
          </tr>
        </thead>
        <tbody>
          <TableBody />
        </tbody>
      </table>
    </div>
  )
}
