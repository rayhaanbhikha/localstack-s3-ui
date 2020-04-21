import React from 'react'
import './styles.css'

export const Bucket = () => {
    return (
        <div className="buckets-table">
            <div className="table-head-container">
                <div class="table-text">
                    <strong className="table-bucket-text">Buckets</strong>&nbsp;&nbsp;<strong className="table-bucket-nums">()</strong>
                </div>
                <div className="table-headings">
                    <div className="table-heading-text">Name</div>
                </div>
            </div>
        </div>
    )
}
