import React from 'react'
import './styles.css'

export const Bucket = () => {
    return (
        <div className="buckets-table">
            <div className="table-head-container">
                <div class="table-text">
                    <strong className="table-bucket-text">Buckets</strong>&nbsp;&nbsp;<strong className="table-bucket-nums">()</strong>
                </div>
            </div>
            <table>
                <div className="table-headings">
                    <tr className="table-column-heading">
                        <th className="table-column-heading-text">Name</th>
                    </tr>
                </div>
                <div className="table-rows">
                    <tr>
                        <td>Some bucket</td>
                    </tr>
                    <tr>
                        <td>Some bucket</td>
                    </tr>
                    <tr>
                        <td>Some bucket</td>
                    </tr>
                </div>
            </table>
        </div>
    )
}
