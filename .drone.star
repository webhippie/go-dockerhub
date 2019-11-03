def main(ctx):
  return [
    testing(ctx),
    changelog(ctx),
    matrix(ctx),
  ]

def testing(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'testing',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'generate',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make generate',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'vet',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make vet',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'staticcheck',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make staticcheck',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'lint',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make lint',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'build',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make build',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'test',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make test',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'codacy',
        'image': 'plugins/codacy:1',
        'pull': 'always',
        'settings': {
          'token': {
            'from_secret': 'codacy_token',
          },
        },
      },
    ],
    'volumes': [
      {
        'name': 'gopath',
        'temp': {},
      },
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
    },
  }

def changelog(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'changelog',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'embedmd',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make embedmd',
        ],
      },
      {
        'name': 'changelog',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make changelog',
        ],
      },
      {
        'name': 'publish',
        'image': 'plugins/git-action:latest',
        'pull': 'always',
        'settings': {
          'actions': [
            'commit',
            'push',
          ],
          'message': 'Automated changelog update [skip ci]',
          'branch': 'master',
          'author_email': 'drone@webhippie.de',
          'author_name': 'Drone',
          'netrc_machine': 'github.com',
          'netrc_username': {
            'from_secret': 'github_username',
          },
          'netrc_password': {
            'from_secret': 'github_password',
          },
        },
        'when': {
          'ref': {
            'exclude': [
              'refs/pull/**',
            ],
          },
        },
      },
    ],
    'depends_on': [
      'testing',
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/pull/**',
      ],
    },
  }

def matrix(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'matrix',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'clone': {
      'disable': True,
    },
    'steps': [
      {
        'name': 'execute',
        'image': 'plugins/matrix',
        'pull': 'always',
        'failure': 'ignore',
        'settings': {
          'username': {
            'from_secret': 'matrix_username',
          },
          'password': {
            'from_secret': 'matrix_password',
          },
          'roomid': {
            'from_secret': 'matrix_roomid',
          },
        },
      },
    ],
    'depends_on': [
      'testing',
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
      ],
      'status': [
        'failure',
      ],
    },
  }
