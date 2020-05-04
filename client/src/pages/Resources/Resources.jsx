import React, { useEffect, useState } from 'react'
import { WrapTable, ResourceRow, BreadCrums } from '../../Components'
import { withRouter } from 'react-router-dom'

export const Component = ({ history }) => {

    const [state, setState] = useState({
        path: "/",
        name: "localstack S3",
        type: "Root",
        breadcrums: [],
        resources: []
    })

    const fetchResources = async (path = "/") => {
        try {
            const res = await fetch(`http://localhost:8080/resource?path=${path}`)
            const data = await res.json();
            setState(prevState => ({
                ...prevState,
                name: data.name,
                path: data.path,
                breadcrums: data.path.split("/").reduce((acc, pathSegment, index) => {
                    const n = acc.length
                    let breadCrum = {}
                    if (index === 0) {
                        breadCrum = {
                            label: "LocalStack S3",
                            url: "/"
                        }
                    } else {
                        breadCrum = {
                            label: pathSegment,
                            url: n > 1 ? acc[n - 1].url + "/" + pathSegment : "/" + pathSegment,
                        }
                    }
                    return [...acc, breadCrum]
                }, []),
                type: data.type,
                resources: data.children || []
            }))
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(() => {
        fetchResources();
    }, [])

    const TableText = () => <>
        <BreadCrums breadcrums={state.breadcrums} fetchResources={fetchResources} />
        {state.type === "Root" ?
            <>
                <strong className="table-bucket-text">Buckets</strong>
            &nbsp;&nbsp;
                <strong className="table-bucket-nums">({state.resources.length})</strong>
            </>
            :
            <strong className="table-bucket-text">{state.name}</strong>
        }
    </>

    const TableBody = () =>
        state.resources.map((resource, index) =>
            <ResourceRow key={`bucketName-${index}`} resource={resource} fetchResources={fetchResources} />
        );

    return WrapTable(TableText, TableBody)
}
export const Resources = withRouter(Component)
