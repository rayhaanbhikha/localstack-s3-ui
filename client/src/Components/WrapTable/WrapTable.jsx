import React from 'react'
import './styles.css'


export const WrapTable = (TableText, TableBody) => {
  return (
    <div className="buckets-table">
      <div className="table-head-container">
        <div className="table-text">
          <TableText />
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
