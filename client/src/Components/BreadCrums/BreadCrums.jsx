import React from 'react';
import { ReactComponent as Chevron } from './chevron.svg'

import { generateBreadCrums } from '../../utils'
import './styles.css'

export const BreadCrums = ({ fetchResources, path }) => {
    let breadcrums = [{
        path: "/",
        name: "Localstack S3"
    }]

    if (path !== "/") {
        breadcrums = [...breadcrums, ...generateBreadCrums(path)]
    }

    return <div className="breadcrums">
        {breadcrums.map((breadcrum, index) =>
            <div key={`${breadcrum.name}-${index}`} className="breadcrum">

                {/* previous breadcrums */}
                {index !== breadcrums.length - 1 &&
                    <>
                        <div onClick={() => fetchResources(breadcrum.path)} className="breadcrum-prev">
                            {breadcrum.name}
                        </div>
                        &nbsp;
                        <Chevron />
                        &nbsp;
                    </>
                }


                {/* last breadcrum */}
                {index === breadcrums.length - 1 &&
                    <div className="breadcrum-current">
                        {breadcrum.name}
                    </div>
                }
            </div>
        )}
    </div>
}