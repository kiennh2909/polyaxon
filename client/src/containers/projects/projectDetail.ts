import * as _ from 'lodash';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { Dispatch } from 'redux';

import * as actions from '../../actions/projects';
import ProjectDetail from '../../components/projects/projectDetail';
import { AppState } from '../../constants/types';
import { getProjectUniqueName } from '../../constants/utils';

export function mapStateToProps(state: AppState, params: any) {
  const projectUniqueName = getProjectUniqueName(
    params.match.params.user,
    params.match.params.projectName);
  return _.includes(state.projects.uniqueNames, projectUniqueName) ?
    {project: state.projects.byUniqueNames[projectUniqueName]} :
    {project: null};
}

export interface DispatchProps {
  onDelete: () => actions.ProjectAction;
  onArchive: () => actions.ProjectAction;
  onRestore: () => actions.ProjectAction;
  stopNotebook: () => actions.ProjectAction;
  stopTensorboard: () => actions.ProjectAction;
  fetchData?: () => actions.ProjectAction;
  bookmark: () => actions.ProjectAction;
  unbookmark: () => actions.ProjectAction;
}

export function mapDispatchToProps(dispatch: Dispatch<actions.ProjectAction>, params: any): DispatchProps {
  return {
    onDelete: () => dispatch(
      actions.deleteProject(
        getProjectUniqueName(
          params.match.params.user,
          params.match.params.projectName),
        true)),
    onArchive: () => dispatch(
      actions.archiveProject(
        getProjectUniqueName(
          params.match.params.user,
          params.match.params.projectName),
        true)),
    onRestore: () => dispatch(
      actions.restoreProject(
        getProjectUniqueName(
          params.match.params.user,
          params.match.params.projectName))),
    stopNotebook: () => dispatch(
      actions.stopNotebook(
        getProjectUniqueName(
          params.match.params.user,
          params.match.params.projectName))),
    stopTensorboard: () => dispatch(
      actions.stopTensorboard(
        getProjectUniqueName(
          params.match.params.user,
          params.match.params.projectName))),
    fetchData: () => dispatch(
      actions.fetchProject(
        params.match.params.user,
        params.match.params.projectName)),
    bookmark: () => dispatch(
      actions.bookmark(getProjectUniqueName(
        params.match.params.user,
        params.match.params.projectName))),
    unbookmark: () => dispatch(
      actions.unbookmark(getProjectUniqueName(
        params.match.params.user,
        params.match.params.projectName))),
  };
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(ProjectDetail));
