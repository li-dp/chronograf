import reducer from 'shared/reducers/annotations'

import {loadAnnotations, updateAnnotation} from 'shared/actions/annotations'

const a1 = {
  id: '1',
  group: '',
  name: 'anno1',
  time: '1515716169000',
  duration: '',
  text: 'you have no swoggels',
}

describe.only('Shared.Reducers.annotations', () => {
  it('can load the annotations', () => {
    const state = []
    const expected = [{time: '0', duration: ''}]
    const actual = reducer(state, loadAnnotations(expected))

    expect(actual).to.deep.equal(expected)
  })

  it('can update an annotation', () => {
    const state = [a1]
    const expected = [{...a1, time: ''}]
    const actual = reducer(state, updateAnnotation(expected[0]))

    expect(actual).to.deep.equal(expected)
  })
})