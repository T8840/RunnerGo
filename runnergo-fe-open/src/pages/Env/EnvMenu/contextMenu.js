import React from 'react';
import ContextMenu from '@components/ContextMenu';
import MenuItem from '@components/ContextMenu/MenuItem';
import { Modal, Message } from 'adesign-react';
import i18next from 'i18next';
import Bus, { useEventBus } from '@utils/eventBus';


export const handleShowContextMenu = (e, item, callback) => {


    const contextMenuRef = React.createRef(null);

    const HoverMenu = (
        <div>
            <ContextMenu
                onItemClick={(action) => {
                    callback && callback(action);
                     contextMenuRef?.current?.hideMenu();
                }}
                style={{ width: 140 }}
                hoverStyle={{ width: 140 }}
            >
                <MenuItem action={1}>
                    { i18next.t('env.editEnv') }
                </MenuItem>
                <MenuItem action={2}>
                    { i18next.t('env.cloneEnv') }
                </MenuItem>
                <MenuItem style={{ color: 'var(--delete-red)' }} action={3}>
                    { i18next.t('env.deleteEnv') }
                </MenuItem>
            </ContextMenu>
        </div>
    );
    Modal.Show(HoverMenu, { x: e.pageX, y: e.pageY }, contextMenuRef);
};
